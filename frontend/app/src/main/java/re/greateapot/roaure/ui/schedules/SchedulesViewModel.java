package re.greateapot.roaure.ui.schedules;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import java.util.List;
import java.util.Objects;

import re.greateapot.roaure.api.RoaureServiceClient;
import re.greateapot.roaure.api.dto.Schedule;
import re.greateapot.roaure.models.StatusWithCallback;

public class SchedulesViewModel extends ViewModel {
    private final MutableLiveData<List<Schedule>> schedulesValue = new MutableLiveData<>();

    // Thx to Flutter BLoC, I know how to pass List updates :)
    private final MutableLiveData<Integer> schedulesVersionValue = new MutableLiveData<>(0);
    private final MutableLiveData<StatusWithCallback> statusValue = new MutableLiveData<>();

    public LiveData<List<Schedule>> getSchedulesValue() {
        return schedulesValue;
    }

    public LiveData<Integer> getSchedulesVersionValue() {
        return schedulesVersionValue;
    }

    public LiveData<StatusWithCallback> getStatusValue() {
        return statusValue;
    }


    public void listSchedules() {
        RoaureServiceClient.getInstance().listSchedules(
                value -> {
                    schedulesValue.postValue(value.getSchedulesList());

                    var version = schedulesVersionValue.getValue();
                    if (version == null) version = 0;
                    schedulesVersionValue.postValue(version + 1);
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(status, this::listSchedules));
                },
                () -> { /* nothing */ }
        );
    }

    public void deleteSchedule(String id) {
        RoaureServiceClient.getInstance().deleteSchedule(
                id,
                value -> {
                    var schedules = schedulesValue.getValue();
                    if (schedules == null) {
                        listSchedules();
                        return;
                    }
                    for (int i = 0; i < schedules.size(); i++) {
                        Schedule schedule = schedules.get(i);
                        if (Objects.equals(schedule.getId(), id)) {
                            schedules.remove(i);
                            break;
                        }
                    }

                    var version = schedulesVersionValue.getValue();
                    if (version == null) version = 0;
                    schedulesVersionValue.postValue(version + 1);
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(status, this::listSchedules));
                },
                () -> { /* nothing */ }
        );
    }

}