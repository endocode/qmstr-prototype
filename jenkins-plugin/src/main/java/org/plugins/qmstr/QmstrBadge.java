package org.plugins.qmstr;

import hudson.Extension;
import hudson.model.BuildBadgeAction;
import net.sf.json.JSONObject;

@Extension
public class QmstrBadge implements BuildBadgeAction{

    public String getStatus() {
        String status = "Qmstr not running";
        QmstrHttpClient qmstr = new QmstrHttpClient("http://localhost:9000");

        JSONObject health = qmstr.health();
        if (health != null) {
            String result = health.getString("running");
            if (result.equals("ok")){
                status = "Qmstr is running";
            }
        }
        return status;
    }

    @Override
    public String getIconFileName() {
        // TODO Auto-generated method stub
        return null;
    }

    @Override
    public String getDisplayName() {
        // TODO Auto-generated method stub
        return "Message";
    }
    @Override
    public String getUrlName() {
        // TODO Auto-generated method stub
        return null;
    }
}
