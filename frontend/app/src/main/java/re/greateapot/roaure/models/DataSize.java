package re.greateapot.roaure.models;

import androidx.annotation.NonNull;

import java.util.Locale;
import java.util.Objects;

public class DataSize {

    private static final DataSize Bit = new DataSize(1.);
    private static final DataSize KBit = new DataSize(1024.);
    private static final DataSize MBit = new DataSize(1024. * 1024);
    private static final DataSize GBit = new DataSize(1024. * 1024 * 1024);
    private static final DataSize Byte = new DataSize(8.);
    private static final DataSize KByte = new DataSize(8. * 1024);
    private static final DataSize MByte = new DataSize(8. * 1024 * 1024);
    private static final DataSize GByte = new DataSize(8. * 1024 * 1024 * 1024);

    private double value;

    public DataSize(double value) {
        this.value = value;
    }

    public double getValue() {
        return value;
    }

    public void setValue(double value) {
        this.value = value;
    }

    @Override
    public boolean equals(Object o) {
        if (o == null || getClass() != o.getClass()) return false;
        DataSize dataSize = (DataSize) o;
        return Double.compare(value, dataSize.value) == 0;
    }

    @Override
    public int hashCode() {
        return Objects.hashCode(value);
    }

    @NonNull
    @Override
    public String toString() {
        var locale = Locale.getDefault();

        if (value >= GBit.value ) {
            return String.format(locale,"%.2f GBit", value / GBit.value);
        } else if (value >= MBit.value ) {
            return String.format(locale,"%.2f MBit", value / MBit.value);
        }else if (value >= KBit.value ) {
            return String.format(locale,"%.2f KBit", value / KBit.value);
        }else {
            return String.format(locale,"%.2f Bit", value / Bit.value);
        }
    }
}
