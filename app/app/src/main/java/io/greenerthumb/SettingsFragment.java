package io.greenerthumb;


import android.os.Bundle;
import android.support.v7.preference.PreferenceFragmentCompat;

/**
 * SettingsFragment binds the settings file to a fragment.
 */
public class SettingsFragment extends PreferenceFragmentCompat {
    @Override
    public void onCreatePreferences(Bundle savedInstanceState, String rootKey) {
        setPreferencesFromResource(R.xml.settings, rootKey);
    }
}