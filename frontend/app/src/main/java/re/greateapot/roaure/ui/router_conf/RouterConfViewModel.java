package re.greateapot.roaure.ui.router_conf;

import androidx.lifecycle.LiveData;
import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;

import re.greateapot.roaure.api.RoaureServiceClient;
import re.greateapot.roaure.api.dto.RouterConf;
import re.greateapot.roaure.models.StatusWithCallback;

public class RouterConfViewModel extends ViewModel {
    private final MutableLiveData<String> hostValue = new MutableLiveData<>();
    private final MutableLiveData<String> usernameValue = new MutableLiveData<>();
    private final MutableLiveData<String> passwordValue = new MutableLiveData<>();
    private final MutableLiveData<StatusWithCallback> statusValue = new MutableLiveData<>();

    public LiveData<String> getHostValue() {
        return hostValue;
    }

    public LiveData<String> getUsernameValue() {
        return usernameValue;
    }

    public LiveData<String> getPasswordValue() {
        return passwordValue;
    }

    public LiveData<StatusWithCallback> getStatusValue() {
        return statusValue;
    }

    public void getConf() {
        RoaureServiceClient.getInstance().getRouterConf(
                value -> {
                    hostValue.postValue(value.getHost());
                    usernameValue.postValue(value.getUsername());
                    passwordValue.postValue(value.getPassword());
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(status, this::getConf));
                },
                () -> { /* nothing */ }
        );
    }

    public void updateConf(String host, String username, String password) {
        RoaureServiceClient.getInstance().updateRouterConf(
                RouterConf.newBuilder().setHost(host).setUsername(username).setPassword(password).build(),
                value -> {
                    hostValue.postValue(host);
                    usernameValue.postValue(username);
                    passwordValue.postValue(password);
                },
                status -> {
                    statusValue.postValue(new StatusWithCallback(status, () -> updateConf(host, username, password)));
                },
                () -> { /* nothing */ }
        );
    }
}