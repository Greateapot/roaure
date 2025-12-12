package re.greateapot.roaure.ui.schedule;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.RadioButton;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.fragment.app.FragmentActivity;
import androidx.lifecycle.ViewModelProvider;
import androidx.navigation.Navigation;

import com.google.android.material.button.MaterialButton;
import com.google.android.material.snackbar.Snackbar;
import com.google.android.material.textfield.TextInputLayout;

import java.util.ArrayList;

import re.greateapot.roaure.R;
import re.greateapot.roaure.api.dto.Weekday;

public class ScheduleFragment extends Fragment {

    private String id;
    private Boolean update;
    private ScheduleViewModel mViewModel;

    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_schedule, container, false);
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        mViewModel = new ViewModelProvider(this).get(ScheduleViewModel.class);

        MaterialButton enabledBtn = view.findViewById(R.id.list_view_schedules_row_toggle);

        enabledBtn.setOnClickListener(view1 -> {
            var value = Boolean.TRUE.equals(mViewModel.getEnabledValue().getValue());
            mViewModel.getEnabledValue().postValue(!value);
        });
        view.findViewById(R.id.reset_button).setOnClickListener(view1 -> unpackArguments());
        view.findViewById(R.id.save_button).setOnClickListener(view1 -> {
            var title = getText(view, R.id.schedule_title, "");

            var startsAtHours = Integer.parseInt(getText(view, R.id.schedule_starts_at_hours, "0"));
            var startsAtMinutes = Integer.parseInt(getText(view, R.id.schedule_starts_at_minutes, "0"));

            var endsAtHours = Integer.parseInt(getText(view, R.id.schedule_ends_at_hours, "0"));
            var endsAtMinutes = Integer.parseInt(getText(view, R.id.schedule_ends_at_minutes, "0"));

            var weekdays = new ArrayList<Weekday>();
            if (((RadioButton) view.findViewById(R.id.schedule_weekdays_monday)).isChecked())
                weekdays.add(Weekday.WEEKDAY_MONDAY);
            if (((RadioButton) view.findViewById(R.id.schedule_weekdays_tuesday)).isChecked())
                weekdays.add(Weekday.WEEKDAY_TUESDAY);
            if (((RadioButton) view.findViewById(R.id.schedule_weekdays_wednesday)).isChecked())
                weekdays.add(Weekday.WEEKDAY_WEDNESDAY);
            if (((RadioButton) view.findViewById(R.id.schedule_weekdays_thursday)).isChecked())
                weekdays.add(Weekday.WEEKDAY_THURSDAY);
            if (((RadioButton) view.findViewById(R.id.schedule_weekdays_friday)).isChecked())
                weekdays.add(Weekday.WEEKDAY_FRIDAY);
            if (((RadioButton) view.findViewById(R.id.schedule_weekdays_saturday)).isChecked())
                weekdays.add(Weekday.WEEKDAY_SATURDAY);
            if (((RadioButton) view.findViewById(R.id.schedule_weekdays_sunday)).isChecked())
                weekdays.add(Weekday.WEEKDAY_SUNDAY);

            var enabled = enabledBtn.isSelected();

            if (update) {
                mViewModel.updateSchedule(id, title, startsAtHours, startsAtMinutes, endsAtHours, endsAtMinutes, weekdays, enabled);
            } else {
                mViewModel.createSchedule(title, startsAtHours, startsAtMinutes, endsAtHours, endsAtMinutes, weekdays, enabled);
            }
        });

        mViewModel.getTitleValue().observe(getViewLifecycleOwner(), value -> setText(view, R.id.schedule_title, value));

        mViewModel.getStartsAtHours().observe(getViewLifecycleOwner(), value -> setText(view, R.id.schedule_starts_at_hours, String.valueOf(value)));
        mViewModel.getStartsAtMinutes().observe(getViewLifecycleOwner(), value -> setText(view, R.id.schedule_starts_at_minutes, String.valueOf(value)));

        mViewModel.getEndsAtHours().observe(getViewLifecycleOwner(), value -> setText(view, R.id.schedule_ends_at_hours, String.valueOf(value)));
        mViewModel.getEndsAtMinutes().observe(getViewLifecycleOwner(), value -> setText(view, R.id.schedule_ends_at_minutes, String.valueOf(value)));

        mViewModel.getWeekdayMondayValue().observe(getViewLifecycleOwner(), value -> ((RadioButton) view.findViewById(R.id.schedule_weekdays_monday)).setChecked(value));
        mViewModel.getWeekdayTuesdayValue().observe(getViewLifecycleOwner(), value -> ((RadioButton) view.findViewById(R.id.schedule_weekdays_tuesday)).setChecked(value));
        mViewModel.getWeekdayWednesdayValue().observe(getViewLifecycleOwner(), value -> ((RadioButton) view.findViewById(R.id.schedule_weekdays_wednesday)).setChecked(value));
        mViewModel.getWeekdayThursdayValue().observe(getViewLifecycleOwner(), value -> ((RadioButton) view.findViewById(R.id.schedule_weekdays_thursday)).setChecked(value));
        mViewModel.getWeekdayFridayValue().observe(getViewLifecycleOwner(), value -> ((RadioButton) view.findViewById(R.id.schedule_weekdays_friday)).setChecked(value));
        mViewModel.getWeekdaySaturdayValue().observe(getViewLifecycleOwner(), value -> ((RadioButton) view.findViewById(R.id.schedule_weekdays_saturday)).setChecked(value));
        mViewModel.getWeekdaySundayValue().observe(getViewLifecycleOwner(), value -> ((RadioButton) view.findViewById(R.id.schedule_weekdays_sunday)).setChecked(value));

        mViewModel.getEnabledValue().observe(getViewLifecycleOwner(), value -> {
            FragmentActivity activity = getActivity();
            if (activity == null) return;

            activity.runOnUiThread(() -> {
                if (value == true) {
                    enabledBtn.setText(R.string.schedule_enabled);
                    enabledBtn.setSelected(true);
                } else {
                    enabledBtn.setText(R.string.schedule_disabled);
                    enabledBtn.setSelected(false);
                }
            });
        });

        mViewModel.getStatusValue().observe(getViewLifecycleOwner(), value -> {
            if (value == null) {
                FragmentActivity activity = getActivity();
                if (activity == null) return;

                Navigation
                        .findNavController(activity, R.id.nav_host_fragment_content_main)
                        .navigateUp();
            } else {
                // TODO: code mapper (unavailable, deadline_exceeded & etc -> err occurred; other -> desc)
                String message = value.status.getCode().toString();
                Snackbar
                        .make(view, message, Snackbar.LENGTH_INDEFINITE)
                        .setAction("Retry", view2 -> value.callback.retry())
                        .show();
            }
        });

        unpackArguments();
    }

    protected void unpackArguments() {
        Bundle arguments = getArguments();
        if (arguments == null) arguments = new Bundle();

        id = arguments.getString("id");
        update = arguments.getBoolean("update", false);

        mViewModel.getTitleValue().postValue(arguments.getString("title", "New Schedule"));

        mViewModel.getStartsAtHours().postValue(arguments.getInt("startsAtHours", 0));
        mViewModel.getStartsAtMinutes().postValue(arguments.getInt("startsAtMinutes", 0));

        mViewModel.getEndsAtHours().postValue(arguments.getInt("endsAtHours", 0));
        mViewModel.getEndsAtMinutes().postValue(arguments.getInt("endsAtMinutes", 0));

        mViewModel.getWeekdayMondayValue().postValue(arguments.getBoolean("weekdayMonday", false));
        mViewModel.getWeekdayTuesdayValue().postValue(arguments.getBoolean("weekdayTuesday", false));
        mViewModel.getWeekdayWednesdayValue().postValue(arguments.getBoolean("weekdayWednesday", false));
        mViewModel.getWeekdayThursdayValue().postValue(arguments.getBoolean("weekdayThursday", false));
        mViewModel.getWeekdayFridayValue().postValue(arguments.getBoolean("weekdayFriday", false));
        mViewModel.getWeekdaySaturdayValue().postValue(arguments.getBoolean("weekdaySaturday", false));
        mViewModel.getWeekdaySundayValue().postValue(arguments.getBoolean("weekdaySunday", false));

        mViewModel.getEnabledValue().postValue(arguments.getBoolean("enabled", false));
    }

    protected String getText(@NonNull View view, int resource, String defaultValue) {
        TextInputLayout v0 = view.findViewById(resource);
        var v1 = v0 != null ? v0.getEditText() : null;
        var v2 = v1 != null ? v1.getText() : null;
        return v2 != null ? v2.toString() : defaultValue;
    }

    protected void setText(@NonNull View view, int resource, String value) {
        TextInputLayout v0 = view.findViewById(resource);
        var v1 = v0 != null ? v0.getEditText() : null;
        if (v1 != null) v1.setText(value);
    }
}