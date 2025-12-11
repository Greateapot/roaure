package re.greateapot.roaure.ui.monitor_conf;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import re.greateapot.roaure.api.RoaureServiceClient;
import re.greateapot.roaure.api.dto.MonitorConf;
import re.greateapot.roaure.api.dto.Time;
import re.greateapot.roaure.models.StatusWithCallback;

public class MonitorConfViewModel extends ViewModel {
    private final MutableLiveData<Double> downloadThresholdValue = new MutableLiveData<>();
    private final MutableLiveData<Integer> pollIntervalHoursValue = new MutableLiveData<>();
    private final MutableLiveData<Integer> pollIntervalMinutesValue = new MutableLiveData<>();
    private final MutableLiveData<Integer> badCountLimitValue = new MutableLiveData<>();
    private final MutableLiveData<StatusWithCallback> statusValue = new MutableLiveData<>();

    public LiveData<Integer> getBadCountLimitValue() {
        return badCountLimitValue;
    }

    public LiveData<Integer> getPollIntervalMinutesValue() {
        return pollIntervalMinutesValue;
    }

    public LiveData<Integer> getPollIntervalHoursValue() {
        return pollIntervalHoursValue;
    }

    public LiveData<Double> getDownloadThresholdValue() {
        return downloadThresholdValue;
    }

    public LiveData<StatusWithCallback> getStatusValue() {
        return statusValue;
    }


    public void getConf() {
        RoaureServiceClient.getInstance().getMonitorConf(
                value -> {
                    downloadThresholdValue.postValue(value.getDownloadThreshold());
                    pollIntervalHoursValue.postValue(value.getPollInterval().getHours());
                    pollIntervalMinutesValue.postValue(value.getPollInterval().getMinutes());
                    badCountLimitValue.postValue(value.getBadCountLimit());
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(status, this::getConf));
                },
                () -> { /* nothing */ }
        );
    }

    public void updateConf(double downloadThreshold, int pollIntervalHours, int pollIntervalMinutes, int badCountLimit) {
        RoaureServiceClient.getInstance().updateMonitorConf(
                MonitorConf
                        .newBuilder()
                        .setDownloadThreshold(downloadThreshold)
                        .setPollInterval(Time
                                .newBuilder()
                                .setHours(pollIntervalHours)
                                .setMinutes(pollIntervalMinutes)
                                .build())
                        .setBadCountLimit(badCountLimit)
                        .build(),
                value -> {
                    downloadThresholdValue.postValue(downloadThreshold);
                    pollIntervalHoursValue.postValue(pollIntervalHours);
                    pollIntervalMinutesValue.postValue(pollIntervalMinutes);
                    badCountLimitValue.postValue(badCountLimit);
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(
                            status,
                            () -> updateConf(downloadThreshold, pollIntervalHours, pollIntervalMinutes, badCountLimit)
                    ));
                },
                () -> { /* nothing */ }
        );
    }
}