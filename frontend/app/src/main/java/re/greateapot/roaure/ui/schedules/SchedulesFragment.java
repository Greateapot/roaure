package re.greateapot.roaure.ui.schedules;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ListView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentActivity;
import androidx.lifecycle.ViewModelProvider;
import androidx.navigation.Navigation;

import java.util.ArrayList;

import re.greateapot.roaure.R;
import re.greateapot.roaure.api.dto.Schedule;

public class SchedulesFragment extends Fragment {

    private SchedulesViewModel mViewModel;
    private SchedulesAdapter adapter;

    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_schedules, container, false);
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        mViewModel = new ViewModelProvider(this).get(SchedulesViewModel.class);

        mViewModel.getSchedulesValue().observe(
                getViewLifecycleOwner(),
                value -> {
                    if (adapter != null) {
                        adapter.setSchedules(value);
                    }
                }
        );

        var context = getContext();
        if (context != null) {
            adapter = new SchedulesAdapter(
                    context,
                    R.id.list_view_schedules_row,
                    new ArrayList<>()
            );

            adapter.setOnEdit(schedule -> {
                FragmentActivity activity = getActivity();
                if (activity != null) {
                    Navigation
                            .findNavController(activity, R.id.nav_host_fragment_content_main)
                            .navigate(R.id.schedule_fragment, packSchedule(schedule));
                }

                return null;
            });

            adapter.setOnDelete(schedule -> {
                mViewModel.deleteSchedule(schedule.getId());
                mViewModel.listSchedules();
                return null;
            });

            view.findViewById(R.id.list_view_schedules_refresh).setOnClickListener(view1 -> {
                mViewModel.listSchedules();
            });
            view.findViewById(R.id.list_view_schedules_add_new).setOnClickListener(view1 -> {
                FragmentActivity activity = getActivity();
                if (activity != null) {
                    Navigation
                            .findNavController(activity, R.id.nav_host_fragment_content_main)
                            .navigate(R.id.schedule_fragment);
                }
            });

            ((ListView)view.findViewById(R.id.list_view_schedules)).setAdapter(adapter);
        }

        mViewModel.listSchedules();
    }

    @Override
    public void onResume() {
        super.onResume();

        mViewModel.listSchedules();
    }

    protected Bundle packSchedule(Schedule schedule) {
        Bundle bundle = new Bundle();
        bundle.putString("id", schedule.getId());
        bundle.putBoolean("update", true);

        bundle.putString("title", schedule.getTitle());

        bundle.putInt("startsAtHours", schedule.getStartsAt().getHours());
        bundle.putInt("startsAtMinutes", schedule.getStartsAt().getMinutes());

        bundle.putInt("endsAtHours", schedule.getEndsAt().getHours());
        bundle.putInt("endsAtMinutes", schedule.getEndsAt().getMinutes());

        schedule.getWeekdaysList().forEach(weekday -> {
            switch (weekday) {
                case WEEKDAY_SUNDAY:
                    bundle.putBoolean("weekdaySunday", true);
                    break;
                case WEEKDAY_MONDAY:
                    bundle.putBoolean("weekdayMonday", true);
                    break;
                case WEEKDAY_TUESDAY:
                    bundle.putBoolean("weekdayTuesday", true);
                    break;
                case WEEKDAY_WEDNESDAY:
                    bundle.putBoolean("weekdayWednesday", true);
                    break;
                case WEEKDAY_THURSDAY:
                    bundle.putBoolean("weekdayThursday", true);
                    break;
                case WEEKDAY_FRIDAY:
                    bundle.putBoolean("weekdayFriday", true);
                    break;
                case WEEKDAY_SATURDAY:
                    bundle.putBoolean("weekdaySaturday", true);
                    break;
            }
        });

        bundle.putBoolean("enabled", schedule.getEnabled());

        return bundle;

    }
}