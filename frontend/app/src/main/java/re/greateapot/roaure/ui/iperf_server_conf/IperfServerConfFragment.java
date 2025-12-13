package re.greateapot.roaure.ui.iperf_server_conf;

import android.os.Bundle;
import android.view.LayoutInflater;
import android.view.View;
import android.view.ViewGroup;

import androidx.annotation.NonNull;
import androidx.annotation.Nullable;
import androidx.fragment.app.Fragment;
import androidx.lifecycle.ViewModelProvider;

import com.google.android.material.snackbar.Snackbar;
import com.google.android.material.textfield.TextInputLayout;

import re.greateapot.roaure.R;

public class IperfServerConfFragment extends Fragment {

    private IperfServerConfViewModel mViewModel;

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

        TextInputLayout h0 = view.findViewById(R.id.text_input_iperf_host);
        TextInputLayout p0 = view.findViewById(R.id.text_input_iperf_port);

        view.findViewById(R.id.reset_button).setOnClickListener(view1 -> mViewModel.getConf());
        view.findViewById(R.id.save_button).setOnClickListener(view1 -> {
            var h1 = h0.getEditText();
            var h2 = h1 != null ? h1.getText() : null;
            var host = h2 != null ? h2.toString() : "";

            var p1 = p0 != null ? p0.getEditText() : null;
            var p2 = p1 != null ? p1.getText() : null;
            var port = Integer.parseInt(p2 != null ? p2.toString() : "0");

            // NOTE: server-side input validation
            mViewModel.updateConf(host, port);
        });

        mViewModel.getHostValue().observe(getViewLifecycleOwner(), value -> {
            var h1 = h0.getEditText();
            if (h1 != null) h1.setText(value);
        });

        mViewModel.getPortValue().observe(getViewLifecycleOwner(), value -> {
            var p1 = p0.getEditText();
            if (p1 != null) p1.setText(String.valueOf(value));
        });

        mViewModel.getStatusValue().observe(getViewLifecycleOwner(), value -> {
            String message = value.status.getCode().toString();
            Snackbar
                    .make(view, message, Snackbar.LENGTH_INDEFINITE)
                    .setAction("Retry", view2 -> value.callback.retry())
                    .show();
        });

        mViewModel.getConf();
    }
}