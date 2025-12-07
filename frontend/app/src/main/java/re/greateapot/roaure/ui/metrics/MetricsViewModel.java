package re.greateapot.roaure.ui.metrics;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import io.grpc.Status;
import re.greateapot.roaure.api.RoaureServiceClient;

public class MetricsViewModel extends ViewModel {

    public static class MetricsViewModelError {

        public interface MetricsViewModelErrorCallback {
            void retry();
        }

        public final Status status;
        public final MetricsViewModelErrorCallback callback;

        public MetricsViewModelError(Status status, MetricsViewModelErrorCallback callback) {
            this.status = status;
            this.callback = callback;
        }
    }

    private final MutableLiveData<Double> downloadSpeedValue = new MutableLiveData<>();
    private final MutableLiveData<Integer> badCountValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> rebootRequiredValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> monitorRunningValue = new MutableLiveData<>();
    private final MutableLiveData<MetricsViewModelError> errorValue = new MutableLiveData<>();

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

    public LiveData<MetricsViewModelError> getErrorValue() {
        return errorValue;
    }

    private boolean isStarted = false;

    public void getMetrics() {
        if (isStarted) return;
        isStarted = true;

        RoaureServiceClient.getInstance().getMetrics(
                10,
                metric -> {
                    downloadSpeedValue.postValue(metric.getDownloadSpeed());
                    badCountValue.postValue(metric.getBadCount());
                    rebootRequiredValue.postValue(metric.getRebootRequired());
                    monitorRunningValue.postValue(metric.getMonitorRunning());
                },
                status -> {
                    isStarted = false;
                    errorValue.postValue(new MetricsViewModelError(status, this::getMetrics));
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
                    errorValue.postValue(new MetricsViewModelError(status, this::toggleMonitor));
                },
                () -> { /* nothing */ }
        );
    }
}
