package org.plugins.qmstr;

import hudson.Extension;
import hudson.model.*;
import org.kohsuke.stapler.DataBoundConstructor;

import net.sf.json.JSONObject;
import org.kohsuke.stapler.StaplerRequest;
import jenkins.model.ParameterizedJobMixIn;

import java.io.IOException;
import java.util.logging.Logger;

public class QuartermasterProperty extends JobProperty<Job<?, ?>>  {

    Process qmstr_master;

    @DataBoundConstructor
    public QuartermasterProperty() {
        /**
         * Start qmstr-master
         */

        try {
            Runtime.getRuntime().exec("cd /qmstr-prototype/qmstr/qmstr-master");
        } catch (IOException e) {
            e.printStackTrace();
        }
        try {
            Runtime.getRuntime().exec("./qmstr-master");
        } catch (IOException e) {
            e.printStackTrace();
        }
        String command = "curl http://localhost:8080/report";
        Process p = null;
        try {
            p = Runtime.getRuntime().exec(command);
        } catch (IOException e) {
            e.printStackTrace();
        }
        LOGGER.info("Qmstr running and getting the report: " + p);
        System.out.print(p);
        //end connection
        try {
            Runtime.getRuntime().exec("curl http://localhost:8080/quit");
        } catch (IOException e) {
            e.printStackTrace();
        }
    }

    public String getName() {
        return qmstr_master.toString();
    }

    @Extension
    public static final class DescriptorImpl extends JobPropertyDescriptor {
        /**
         * Used to hide property configuration under checkbox,
         * as of not each job is running with Qmstr build environment
         */
        public static final String QMSTRQ_PROJECT_BLOCK_NAME = "quartermasterProject";

        public boolean isApplicable(Class<? extends Job> jobType) {
            return ParameterizedJobMixIn.ParameterizedJob.class.isAssignableFrom(jobType);
        }

        public String getDisplayName() {
            return QMSTRQ_PROJECT_BLOCK_NAME;
        }

        @Override
        public JobProperty<?> newInstance(StaplerRequest req, JSONObject formData) throws FormException {
            QuartermasterProperty tpp = req.bindJSON(
                    QuartermasterProperty.class,
                    formData.getJSONObject(QMSTRQ_PROJECT_BLOCK_NAME)
            );

            if (tpp == null) {
                LOGGER.fine("Couldn't bind JSON");
                return null;
            }

            return tpp;
        }

    }
    private static final Logger LOGGER = Logger.getLogger(QuartermasterProperty.class.getName());
}

