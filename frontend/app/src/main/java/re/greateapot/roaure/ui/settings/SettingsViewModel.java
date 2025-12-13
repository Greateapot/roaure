package re.greateapot.roaure.ui.settings;

import androidx.lifecycle.MutableLiveData;
import androidx.lifecycle.ViewModel;


public class SettingsViewModel extends ViewModel {
    private final MutableLiveData<String> hostValue = new MutableLiveData<>();
    private final MutableLiveData<Integer> portValue = new MutableLiveData<>();

    public MutableLiveData<String> getHostValue() {
        return hostValue;
    }

    public MutableLiveData<Integer> getPortValue() {
        return portValue;
    }
}