package re.greateapot.roaure.ui.schedules;

import android.annotation.SuppressLint;
import android.content.Context;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;
import android.widget.ArrayAdapter;
import android.widget.TextView;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;

import com.google.android.material.button.MaterialButton;

import java.util.ArrayList;
import java.util.List;
import java.util.Locale;
import java.util.function.Function;

import re.greateapot.roaure.R;
import re.greateapot.roaure.api.dto.Schedule;
import re.greateapot.roaure.api.dto.Weekday;

public class SchedulesAdapter extends ArrayAdapter<Schedule> {

    private Function<Schedule, Boolean> onEdit;
    private Function<Schedule, Boolean> onDelete;
    private List<Schedule> schedules;

    public SchedulesAdapter(@NonNull Context context, int resource, List<Schedule> schedules) {
        super(context, resource, schedules);
        setSchedules(schedules);
    }

    public void setOnEdit(Function<Schedule, Boolean> onEdit) {
        this.onEdit = onEdit;
    }

    public void setOnDelete(Function<Schedule, Boolean> onDelete) {
        this.onDelete = onDelete;
    }

    public void setSchedules(List<Schedule> schedules) {
        this.schedules = schedules;
        this.notifyDataSetChanged();
    }

    @Override
    public int getCount() {
        return schedules.size();
    }

    @Nullable
    @Override
    public Schedule getItem(int position) {
        return schedules.get(position);
    }

    @Override
    public long getItemId(int position) {
        return super.getItemId(position);
    }

    @SuppressLint("InflateParams")
    @NonNull
    @Override
    public View getView(int position, @Nullable View convertView, @NonNull ViewGroup parent) {
        View view = convertView != null
                ? convertView
                : LayoutInflater.from(getContext()).inflate(R.layout.list_view_schedules_row, null);

        Schedule schedule = getItem(position);
        if (schedule == null) return view;

        TextView tvTitle = view.findViewById(R.id.list_view_schedules_row_title);
        TextView tvInterval = view.findViewById(R.id.list_view_schedules_row_interval);
        TextView tvWeekdays = view.findViewById(R.id.list_view_schedules_row_weekdays);
        MaterialButton btnToggle = view.findViewById(R.id.list_view_schedules_row_toggle);
        MaterialButton btnEdit = view.findViewById(R.id.list_view_schedules_row_edit);
        MaterialButton btnDelete = view.findViewById(R.id.list_view_schedules_row_delete);

        tvTitle.setText(schedule.getTitle());
        tvInterval.setText(String.format(
                Locale.getDefault(),
                "Interval: %02d:%02d - %02d:%02d",
                schedule.getStartsAt().getHours(),
                schedule.getStartsAt().getMinutes(),
                schedule.getEndsAt().getHours(),
                schedule.getEndsAt().getMinutes()
        ));

        var weekdays = new ArrayList<String>();
        schedule.getWeekdaysList().forEach(weekday -> {
            weekdays.add(weekdayToString(weekday));
        });
        tvWeekdays.setText(String.format("Weekdays: %s", String.join(", ", weekdays)));

        if (schedule.getEnabled()) {
            btnToggle.setText(R.string.schedule_enabled);
            btnToggle.setSelected(true);
        } else {
            btnToggle.setText(R.string.schedule_disabled);
            btnToggle.setSelected(false);
        }

        btnEdit.setOnClickListener((view1) -> {
            if (onEdit != null) onEdit.apply(schedule);
        });
        btnDelete.setOnClickListener((view1) -> {
            if (onDelete != null) onDelete.apply(schedule);
        });

        return view;
    }

    protected String weekdayToString(Weekday weekday) {
        switch (weekday) {
            case WEEKDAY_SUNDAY:
                return "sunday";
            case WEEKDAY_MONDAY:
                return "monday";
            case WEEKDAY_TUESDAY:
                return "tuesday";
            case WEEKDAY_WEDNESDAY:
                return "wednesday";
            case WEEKDAY_THURSDAY:
                return "thursday";
            case WEEKDAY_FRIDAY:
                return "friday";
            default: // WEEKDAY_SATURDAY
                return "saturday";
        }
    }
}
