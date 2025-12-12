package re.greateapot.roaure.ui.schedule;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import re.greateapot.roaure.api.RoaureServiceClient;
import re.greateapot.roaure.api.dto.Schedule;
import re.greateapot.roaure.api.dto.Time;
import re.greateapot.roaure.api.dto.Weekday;
import re.greateapot.roaure.models.StatusWithCallback;

public class ScheduleViewModel extends ViewModel {
    private final MutableLiveData<String> titleValue = new MutableLiveData<>();

    private final MutableLiveData<Integer> startsAtHours = new MutableLiveData<>();
    private final MutableLiveData<Integer> startsAtMinutes = new MutableLiveData<>();

    private final MutableLiveData<Integer> endsAtHours = new MutableLiveData<>();
    private final MutableLiveData<Integer> endsAtMinutes = new MutableLiveData<>();

    private final MutableLiveData<Boolean> weekdayMondayValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> weekdayTuesdayValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> weekdayWednesdayValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> weekdayThursdayValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> weekdayFridayValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> weekdaySaturdayValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> weekdaySundayValue = new MutableLiveData<>();

    private final MutableLiveData<Boolean> enabledValue = new MutableLiveData<>();

    private final MutableLiveData<StatusWithCallback> statusValue = new MutableLiveData<>();

    public MutableLiveData<Boolean> getEnabledValue() {
        return enabledValue;
    }

    public MutableLiveData<String> getTitleValue() {
        return titleValue;
    }

    public MutableLiveData<Integer> getStartsAtHours() {
        return startsAtHours;
    }

    public MutableLiveData<Integer> getStartsAtMinutes() {
        return startsAtMinutes;
    }

    public MutableLiveData<Integer> getEndsAtHours() {
        return endsAtHours;
    }

    public MutableLiveData<Integer> getEndsAtMinutes() {
        return endsAtMinutes;
    }

    public MutableLiveData<Boolean> getWeekdayMondayValue() {
        return weekdayMondayValue;
    }

    public MutableLiveData<Boolean> getWeekdayTuesdayValue() {
        return weekdayTuesdayValue;
    }

    public MutableLiveData<Boolean> getWeekdayWednesdayValue() {
        return weekdayWednesdayValue;
    }

    public MutableLiveData<Boolean> getWeekdayThursdayValue() {
        return weekdayThursdayValue;
    }

    public MutableLiveData<Boolean> getWeekdayFridayValue() {
        return weekdayFridayValue;
    }

    public MutableLiveData<Boolean> getWeekdaySaturdayValue() {
        return weekdaySaturdayValue;
    }

    public MutableLiveData<Boolean> getWeekdaySundayValue() {
        return weekdaySundayValue;
    }

    public LiveData<StatusWithCallback> getStatusValue() {
        return statusValue;
    }

    public void createSchedule(String title,
                               int startsAtHours,
                               int startsAtMinutes,
                               int endsAtHours,
                               int endsAtMinutes,
                               Iterable<Weekday> weekdays,
                               boolean enabled) {
        RoaureServiceClient.getInstance().createSchedule(
                Schedule
                        .newBuilder()
                        .setEnabled(true)
                        .setTitle(title)
                        .setStartsAt(Time
                                .newBuilder()
                                .setHours(startsAtHours)
                                .setMinutes(startsAtMinutes)
                                .build())
                        .setEndsAt(Time
                                .newBuilder()
                                .setHours(endsAtHours)
                                .setMinutes(endsAtMinutes)
                                .build())
                        .addAllWeekdays(weekdays)
                        .setEnabled(enabled)
                        .build(),
                value -> statusValue.postValue(null),
                status -> statusValue.postValue(new StatusWithCallback(status, () -> createSchedule
                        (title, startsAtHours, startsAtMinutes, endsAtHours, endsAtMinutes, weekdays, enabled))),
                () -> { /* nothing */ }
        );
    }

    public void updateSchedule(String id,
                               String title,
                               int startsAtHours,
                               int startsAtMinutes,
                               int endsAtHours,
                               int endsAtMinutes,
                               Iterable<Weekday> weekdays,
                               boolean enabled) {
        RoaureServiceClient.getInstance().updateSchedule(
                id,
                Schedule
                        .newBuilder()
                        .setEnabled(enabled)
                        .setTitle(title)
                        .setStartsAt(Time
                                .newBuilder()
                                .setHours(startsAtHours)
                                .setMinutes(startsAtMinutes)
                                .build())
                        .setEndsAt(Time
                                .newBuilder()
                                .setHours(endsAtHours)
                                .setMinutes(endsAtMinutes)
                                .build())
                        .addAllWeekdays(weekdays)
                        .build(),
                value -> statusValue.postValue(null),
                status -> statusValue.postValue(new StatusWithCallback(status, () -> updateSchedule
                        (id, title, startsAtHours, startsAtMinutes, endsAtHours, endsAtMinutes, weekdays, enabled))),
                () -> { /* nothing */ }
        );
    }
}