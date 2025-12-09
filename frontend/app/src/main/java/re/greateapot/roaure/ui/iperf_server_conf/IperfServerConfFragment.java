package re.greateapot.roaure.ui.iperf_server_conf;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.lifecycle.ViewModelProvider;

import re.greateapot.roaure.R;

public class IperfServerConfFragment extends Fragment {

    private IperfServerConfViewModel mViewModel;

    public static IperfServerConfFragment newInstance() {
        return new IperfServerConfFragment();
    }

    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_iperf_server_conf, container, false);
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        mViewModel = new ViewModelProvider(this).get(IperfServerConfViewModel.class);
        // TODO: Use the ViewModel
    }

}