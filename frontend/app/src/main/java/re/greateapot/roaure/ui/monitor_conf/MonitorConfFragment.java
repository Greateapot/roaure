package re.greateapot.roaure.ui.monitor_conf;

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

public class MonitorConfFragment extends Fragment {

    private MonitorConfViewModel mViewModel;

    @Override
    public View onCreateView(@NonNull LayoutInflater inflater,
                             @Nullable ViewGroup container,
                             @Nullable Bundle savedInstanceState) {
        return inflater.inflate(R.layout.fragment_monitor_conf, container, false);
    }

    @Override
    public void onViewCreated(@NonNull View view, @Nullable Bundle savedInstanceState) {
        super.onViewCreated(view, savedInstanceState);
        mViewModel = new ViewModelProvider(this).get(MonitorConfViewModel.class);

        TextInputLayout dt0 = view.findViewById(R.id.text_input_monitor_download_threshold);
        TextInputLayout pih0 = view.findViewById(R.id.text_input_monitor_poll_interval_hours);
        TextInputLayout pim0 = view.findViewById(R.id.text_input_monitor_poll_interval_minutes);
        TextInputLayout bcl0 = view.findViewById(R.id.text_input_monitor_bad_count_limit);

        view.findViewById(R.id.reset_button).setOnClickListener(view1 -> mViewModel.getConf());
        view.findViewById(R.id.save_button).setOnClickListener(view1 -> {
            var dt1 = dt0.getEditText();
            var dt2 = dt1 != null ? dt1.getText() : null;
            var downloadThreshold = Double.parseDouble(dt2 != null ? dt2.toString() : "0.0");

            var pih1 = pih0 != null ? pih0.getEditText() : null;
            var pih2 = pih1 != null ? pih1.getText() : null;
            var pollIntervalHours = Integer.parseInt(pih2 != null ? pih2.toString() : "0");

            var pim1 = pim0 != null ? pim0.getEditText() : null;
            var pim2 = pim1 != null ? pim1.getText() : null;
            var pollIntervalMinutes = Integer.parseInt(pim2 != null ? pim2.toString() : "0");

            var bcl1 = bcl0 != null ? bcl0.getEditText() : null;
            var bcl2 = bcl1 != null ? bcl1.getText() : null;
            var badCountLimit = Integer.parseInt(bcl2 != null ? bcl2.toString() : "0");

            // NOTE: server-side input validation
            mViewModel.updateConf(downloadThreshold, pollIntervalHours, pollIntervalMinutes, badCountLimit);
        });

        mViewModel.getDownloadThresholdValue().observe(getViewLifecycleOwner(), value -> {
            var dt1 = dt0.getEditText();
            if (dt1 != null) dt1.setText(String.valueOf(value));
        });

        mViewModel.getPollIntervalHoursValue().observe(getViewLifecycleOwner(), value -> {
            var pih1 = pih0.getEditText();
            if (pih1 != null) pih1.setText(String.valueOf(value));
        });

        mViewModel.getPollIntervalMinutesValue().observe(getViewLifecycleOwner(), value -> {
            var pim1 = pim0.getEditText();
            if (pim1 != null) pim1.setText(String.valueOf(value));
        });

        mViewModel.getBadCountLimitValue().observe(getViewLifecycleOwner(), value -> {
            var bcl1 = bcl0.getEditText();
            if (bcl1 != null) bcl1.setText(String.valueOf(value));
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