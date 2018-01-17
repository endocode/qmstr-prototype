package org.plugins.qmstr;

import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

import org.kohsuke.stapler.DataBoundConstructor;

import hudson.Extension;
import hudson.Launcher;
import hudson.model.AbstractBuild;
import hudson.model.AbstractProject;
import hudson.model.BuildListener;
import hudson.tasks.BuildStepDescriptor;
import hudson.tasks.BuildStepMonitor;
import hudson.tasks.Publisher;
import hudson.tasks.Recorder;
import net.sf.json.JSONArray;
import net.sf.json.JSONObject;

@Extension
public class QmstrReport extends Recorder {

    @DataBoundConstructor
    public QmstrReport(){

    }

    @Override
    public BuildStepMonitor getRequiredMonitorService() {
        return BuildStepMonitor.NONE;
    }

    @Override
    public boolean perform(AbstractBuild<?, ?> build, Launcher launcher, BuildListener listener) throws InterruptedException, IOException {

        QmstrHttpClient qmstr = new QmstrHttpClient("http://localhost:9000");

        JSONObject linkedTargets = qmstr.linkedTargets();
        if (linkedTargets == null) {
            return false;
        }

        Map<String, String> map = new HashMap<>();
        if (!linkedTargets.has("linkedtargets")) {

            return true;
        }

        JSONArray linkedtargetsArray = linkedTargets.getJSONArray("linkedtargets");

        for (int i=0; i< linkedtargetsArray.size(); i++){

            String targetName = linkedtargetsArray.get(i).toString();
            JSONObject reporttargetNameSpecific  = qmstr.report(targetName);

            map.put(targetName, reporttargetNameSpecific.getString("report"));
        }

        build.addAction(new BuildReportAction(map));
        
        return true;
    }

    @Override
    public BuildStepDescriptor getDescriptor() {
        return (DescriptorImpl) super.getDescriptor();
    }

    @Extension
    public static class DescriptorImpl extends BuildStepDescriptor<Publisher> {

        public DescriptorImpl() {
            super();
        }

        public DescriptorImpl(Class<? extends Publisher> clazz) {
            super(clazz);
            // TODO Auto-generated constructor stub
        }

        @Override
        public boolean isApplicable(Class<? extends AbstractProject> jobType) {
            return true;
        }

        @Override
        public String getDisplayName() {
            return "Generate reuse badge";
        }

    }
}
