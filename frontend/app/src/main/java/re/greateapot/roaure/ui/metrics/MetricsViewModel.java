package re.greateapot.roaure.ui.metrics;

import android.util.Log;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import re.greateapot.roaure.api.RoaureServiceClient;

public class MetricsViewModel extends ViewModel {
    private static final String TAG = "MetricsVM";

    private final MutableLiveData<Boolean> isStarted = new MutableLiveData<>(false);
    private final MutableLiveData<Double> metricValue = new MutableLiveData<>(0.0);

    public LiveData<Double> getMetricValue() {
        return metricValue;
    }

    public void getMetrics() {
        if (Boolean.TRUE.equals(isStarted.getValue())) return;
        isStarted.postValue(true);

        RoaureServiceClient.getInstance().getMetrics(
                10,
                metric -> {
                    Log.i(TAG, String.format("DLS: %f", metric.getDownloadSpeed()));
                    metricValue.postValue(metric.getDownloadSpeed());
                },
                status -> {
                    // TODO: toast
                    Log.i(TAG, String.format("Error (%s): %s", status.getCode().toString(), status.getDescription()));
                },
                () -> {}
        );
    }
}
