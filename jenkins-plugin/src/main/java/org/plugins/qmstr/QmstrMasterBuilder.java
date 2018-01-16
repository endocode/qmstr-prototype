package org.plugins.qmstr;

import java.io.IOException;
import java.util.ArrayList;
import java.util.StringTokenizer;

import org.kohsuke.stapler.DataBoundConstructor;

import hudson.Extension;
import hudson.Launcher;
import hudson.model.AbstractBuild;
import hudson.model.AbstractProject;
import hudson.model.BuildListener;
import hudson.tasks.BuildStepDescriptor;
import hudson.tasks.Builder;


public class QmstrMasterBuilder extends Builder {

    @DataBoundConstructor
    public QmstrMasterBuilder(){
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

        QmstrHttpClient client = new QmstrHttpClient("http://localhost:9000");
        client.quit();

        QuartermasterProperty prop = build.getProject().getProperty(QuartermasterProperty.class);
        String pathToQMstrMaster;
        //Process process;

        if (prop != null){
            pathToQMstrMaster = prop.getPath();
        } else {
            return false;
        }

        //launcher.launch(pathToQMstrMaster, build.getEnvVars(), listener.getLogger(),build.getProject().getWorkspace());
        // process = Runtime.getRuntime().exec(pathToQMstrMaster);

        StringTokenizer st = new StringTokenizer(pathToQMstrMaster);
        ArrayList<String> list = new ArrayList<>();
        while (st.hasMoreTokens()) {
            list.add(st.nextToken());
        }

        ProcessBuilder pb = new ProcessBuilder(list).redirectOutput(ProcessBuilder.Redirect.INHERIT).redirectError(ProcessBuilder.Redirect.INHERIT);
        pb.start();
        
        
        // Check if the master is actually running first
        build.addAction(new QmstrBadge());
        
        return true;
    }

}
