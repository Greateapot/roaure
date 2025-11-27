package re.greateapot.roaure.ui.metrics;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import re.greateapot.roaure.api.RoaureServiceClient;

public class MetricsViewModel extends ViewModel {

    private final MutableLiveData<Boolean> isStarted = new MutableLiveData<>(false);
    private final MutableLiveData<Double> metricValue = new MutableLiveData<>();
    private final MutableLiveData<Integer> badCountValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> rebootRequiredValue = new MutableLiveData<>();
    private final MutableLiveData<Boolean> monitorRunningValue = new MutableLiveData<>();

    public LiveData<Double> getMetricValue() {
        return metricValue;
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


    public void getMetrics() {
        if (Boolean.TRUE.equals(isStarted.getValue())) return;
        isStarted.postValue(true);

        RoaureServiceClient.getInstance().getMetrics(
                10,
                metric -> {
                    metricValue.postValue(metric.getDownloadSpeed());
                    badCountValue.postValue(metric.getBadCount());
                    rebootRequiredValue.postValue(metric.getRebootRequired());
                    monitorRunningValue.postValue(metric.getMonitorRunning());
                },
                status -> {
                    // TODO: show snackbar with retry button
                    isStarted.postValue(false);
                },
                () -> {
                    isStarted.postValue(false);
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
