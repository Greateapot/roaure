package re.greateapot.roaure.ui.iperf_server_conf;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import re.greateapot.roaure.api.RoaureServiceClient;
import re.greateapot.roaure.api.dto.IperfServerConf;
import re.greateapot.roaure.models.StatusWithCallback;


public class IperfServerConfViewModel extends ViewModel {
    private final MutableLiveData<String> hostValue = new MutableLiveData<>();
    private final MutableLiveData<Integer> portValue = new MutableLiveData<>();
    private final MutableLiveData<StatusWithCallback> statusValue = new MutableLiveData<>();

    public LiveData<String> getHostValue() {
        return hostValue;
    }

    public LiveData<Integer> getPortValue() {
        return portValue;
    }

    public LiveData<StatusWithCallback> getStatusValue() {
        return statusValue;
    }

    public void getConf() {
        RoaureServiceClient.getInstance().getIperfServerConf(
                value -> {
                    hostValue.postValue(value.getHost());
                    portValue.postValue(value.getPort());
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(status, this::getConf));
                },
                () -> { /* nothing */ }
        );
    }

    public void updateConf(String host, int port) {
        RoaureServiceClient.getInstance().updateIperfServerConf(
                IperfServerConf.newBuilder().setHost(host).setPort(port).build(),
                value -> {
                    hostValue.postValue(host);
                    portValue.postValue(port);
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(status, () -> updateConf(host, port)));
                },
                () -> { /* nothing */ }
        );
    }
}