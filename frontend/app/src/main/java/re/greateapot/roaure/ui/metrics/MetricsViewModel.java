package re.greateapot.roaure.ui.metrics;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import re.greateapot.roaure.api.RoaureServiceClient;
import re.greateapot.roaure.models.DataSize;
import re.greateapot.roaure.models.StatusWithCallback;

public class MetricsViewModel extends ViewModel {

    private final MutableLiveData<DataSize> downloadSpeedValue = new MutableLiveData<>();
    private final MutableLiveData<Integer> badCountValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> rebootRequiredValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> monitorRunningValue = new MutableLiveData<>();
    private final MutableLiveData<StatusWithCallback> statusValue = new MutableLiveData<>();

    public LiveData<DataSize> getDownloadSpeedValue() {
        return downloadSpeedValue;
    }

    public LiveData<Integer> getBadCountValue() {
        return badCountValue;
    }

    public LiveData<Boolean> getRebootRequiredValue() {
        return rebootRequiredValue;
    }

    public LiveData<Boolean> getMonitorRunningValue() {
        return monitorRunningValue;
    }

    public LiveData<StatusWithCallback> getStatusValue() {
        return statusValue;
    }

    private boolean isStarted = false;

    public void getMetrics() {
        if (isStarted) return;
        isStarted = true;

        RoaureServiceClient.getInstance().getMetrics(
                10,
                metric -> {
                    downloadSpeedValue.postValue(new DataSize(metric.getDownloadSpeed()));
                    badCountValue.postValue(metric.getBadCount());
                    rebootRequiredValue.postValue(metric.getRebootRequired());
                    monitorRunningValue.postValue(metric.getMonitorRunning());
                },
                status -> {
                    isStarted = false;
                    statusValue.postValue(new StatusWithCallback(status, this::getMetrics));
                },
                () -> {
                    // There's nothing we can do...
                    isStarted = false;
                }
        );
    }

    public void toggleMonitor() {
        RoaureServiceClient.getInstance().toggleMonitor(
                e -> {
                    boolean monitorRunning = Boolean.TRUE.equals(monitorRunningValue.getValue());
                    monitorRunningValue.postValue(!monitorRunning);
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(status, this::toggleMonitor));
                },
                () -> { /* nothing */ }
        );
    }

    public void reboot() {
        RoaureServiceClient.getInstance().reboot(
                e -> {
                    statusValue.postValue(null);
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(status, this::toggleMonitor));
                },
                () -> { /* nothing */ }
        );
    }
}
