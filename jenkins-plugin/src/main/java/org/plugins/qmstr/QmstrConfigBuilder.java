package org.plugins.qmstr;

import hudson.FilePath;
import hudson.Extension;
import hudson.Launcher;
import hudson.model.AbstractBuild;
import hudson.model.AbstractProject;
import hudson.model.BuildListener;
import hudson.tasks.BuildStepDescriptor;
import hudson.tasks.Builder;
import net.sf.json.JSONObject;

import org.kohsuke.stapler.DataBoundConstructor;

import java.io.IOException;
import java.time.LocalDateTime;


public class QmstrConfigBuilder extends Builder {

    @DataBoundConstructor
    public QmstrConfigBuilder(){
    }

    @Extension
    public static class Descriptor extends BuildStepDescriptor<Builder> {

        @Override
        public boolean isApplicable(Class<? extends AbstractProject> jobType) {
            return true;
        }
        @Override
        public String getDisplayName() {
            return "execute Qmstr-master server";
        }
    }

    @Override
    public boolean perform(AbstractBuild<?, ?> build, Launcher launcher, BuildListener listener) throws InterruptedException, IOException {
        FilePath wd = build.getWorkspace();
        JSONObject configData = new JSONObject();
        configData.put("workdir", wd.absolutize().toString());

        QmstrHttpClient client = new QmstrHttpClient("http://localhiost:9000");
        client.configure(configData);

        LocalDateTime now = LocalDateTime.now();
        LocalDateTime timeout = now.plusMinutes(10);

        while (now.isBefore(timeout)) {
            JSONObject health = client.health();
            if (health.has("scanned")) {
                if (health.getBoolean("scanned")) {
                    return true;
                }
            } else {
                // qmstr does not support this
                return false;
            }
            now = LocalDateTime.now();
        }
        return false;
    }

}
