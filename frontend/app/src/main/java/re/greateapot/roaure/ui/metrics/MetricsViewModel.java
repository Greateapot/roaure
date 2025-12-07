package re.greateapot.roaure.ui.metrics;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import io.grpc.Status;
import re.greateapot.roaure.api.RoaureServiceClient;

public class MetricsViewModel extends ViewModel {

    private final MutableLiveData<Double> downloadSpeedValue = new MutableLiveData<>();
    private final MutableLiveData<Integer> badCountValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> rebootRequiredValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> monitorRunningValue = new MutableLiveData<>();

    public LiveData<Double> getDownloadSpeedValue() {
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

    private boolean isStarted = false;

    public void getMetrics() {
        if (isStarted) return;
        isStarted =true;

        RoaureServiceClient.getInstance().getMetrics(
                10,
                metric -> {
                    downloadSpeedValue.postValue(metric.getDownloadSpeed());
                    badCountValue.postValue(metric.getBadCount());
                    rebootRequiredValue.postValue(metric.getRebootRequired());
                    monitorRunningValue.postValue(metric.getMonitorRunning());
                },
                status -> {
                    // TODO: show snackbar with retry button
                    isStarted = false;
                },
                () -> {
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
                    // TODO: show snackbar with retry button
                },
                () -> { /* nothing */ }
        );
    }
}
