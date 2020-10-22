package com.huawei.hms.cordova.push.remote;

import android.content.Context;
import android.content.SharedPreferences;
import android.util.Log;
import android.webkit.WebView;
import android.webkit.WebViewClient;

import com.huawei.hms.cordova.push.HMSPush;
import com.huawei.hms.cordova.push.hmslogger.HMSLogger;
import com.huawei.hms.cordova.push.local.HmsLocalNotification;
import com.huawei.hms.cordova.push.utils.ApplicationUtils;
import com.huawei.hms.cordova.push.utils.RemoteMessageUtils;
import com.huawei.hms.push.HmsMessageService;
import com.huawei.hms.push.RemoteMessage;
import com.huawei.hms.push.SendException;

import org.json.JSONArray;
import org.json.JSONException;
import org.json.JSONObject;

import java.util.Iterator;

public class HmsPushMessageService extends HmsMessageService {
    private final String TAG = HmsPushMessageService.class.getSimpleName();
    private HMSLogger hmsLogger;
    private WebView webView;

    @Override
    public void onMessageReceived(RemoteMessage message) {
        Log.w(TAG, "** onMessageReceived **");
        hmsLogger = HMSLogger.getInstance(getApplicationContext());
        try {
            boolean isApplicationInForeground = ApplicationUtils.isApplicationInForeground(getApplicationContext());
            if (isApplicationInForeground) {
                HmsMessagePublisher.sendMessageReceivedEvent(message);
                hmsLogger.sendPeriodicEvent("onMessageReceived");
            } else {
                String myAppId = getApplicationContext().getApplicationInfo().uid + "";
                SharedPreferences sharedPref = getApplication().
                        getApplicationContext().
                        getSharedPreferences(getPackageName() + "." + myAppId, Context.MODE_PRIVATE);
                if (webView == null) {
                    webView = new WebView(getApplicationContext());
                    webView.setWebViewClient(new WebViewClient());
                    webView.getSettings().setJavaScriptEnabled(true);
                    webView.addJavascriptInterface(new BackgroundJavaScriptInterface(getApplication()), "HmsLocalNotification");

                }
                String preFunction = sharedPref.getString("data", null);
                if (preFunction != null)
                    preFunction = preFunction.replace("=>", "");
                String function = String.format("function callback%s", preFunction);
                webView.loadUrl("javascript:" + function);
                webView.loadUrl(String.format("javascript:callback(%s);", RemoteMessageUtils.fromMap(message)));
            }
        } catch (JSONException e) {
            Log.w(TAG, "onMessageReceived: " + e.getLocalizedMessage());
            hmsLogger.sendPeriodicEvent("onMessageReceived", e.getLocalizedMessage());
        }
    }

    @Override
    public void onDeletedMessages() {
        if (HMSPush.getPlugin() == null)
            return;
        hmsLogger = HMSLogger.getInstance(HMSPush.getPlugin().cordova.getContext());
        Log.w(TAG, "** onDeletedMessages **");
        hmsLogger.sendPeriodicEvent("onDeletedMessages");
    }

    @Override
    public void onMessageSent(String msgId) {
        if (HMSPush.getPlugin() == null)
            return;
        hmsLogger = HMSLogger.getInstance(HMSPush.getPlugin().cordova.getContext());
        Log.w(TAG, "** onMessageSent **");
        try {
            HmsMessagePublisher.sendOnMessageSentEvent(msgId);
            hmsLogger.sendPeriodicEvent("onMessageSent");
        } catch (JSONException e) {
            Log.w(TAG, "onMessageSent: " + e.getLocalizedMessage());
            hmsLogger.sendPeriodicEvent("onMessageSent", e.getLocalizedMessage());
        }
    }

    @Override
    public void onSendError(String msgId, Exception exception) {
        if (HMSPush.getPlugin() == null)
            return;
        hmsLogger = HMSLogger.getInstance(HMSPush.getPlugin().cordova.getContext());
        Log.w(TAG, "** onSendError **");
        int errorCode = ((SendException) exception).getErrorCode();
        String errorInfo = exception.getLocalizedMessage();
        try {
            HmsMessagePublisher.sendOnMessageSentErrorEvent(msgId, errorCode, errorInfo);
            hmsLogger.sendPeriodicEvent("onSendError");
        } catch (JSONException e) {
            Log.w(TAG, "onSendError: " + e.getLocalizedMessage());
            hmsLogger.sendPeriodicEvent("onSendError", e.getLocalizedMessage());
        }
    }

    @Override
    public void onMessageDelivered(String msgId, Exception e) {
        if (HMSPush.getPlugin() == null)
            return;
        hmsLogger = HMSLogger.getInstance(HMSPush.getPlugin().cordova.getContext());
        Log.w(TAG, "** onMessageDelivered **");
        try {
            if (e == null) {
                HmsMessagePublisher.sendOnMessageDeliveredEvent(msgId, 0, "");
            } else {
                HmsMessagePublisher.sendOnMessageDeliveredEvent(msgId, ((SendException) e).getErrorCode(), e.getLocalizedMessage());
            }
            hmsLogger.sendPeriodicEvent("onMessageDelivered");
        } catch (JSONException ex) {
            hmsLogger.sendPeriodicEvent("onMessageDelivered", ex.getLocalizedMessage());
        }
    }

    @Override
    public void onTokenError(Exception e) {
        if (HMSPush.getPlugin() == null)
            return;
        hmsLogger = HMSLogger.getInstance(HMSPush.getPlugin().cordova.getContext());
        Log.w(TAG, "** onTokenError **");
        try {
            HmsMessagePublisher.sendTokenErrorEvent(e);
            hmsLogger.sendPeriodicEvent("onTokenError");
        } catch (JSONException ex) {
            Log.w(TAG, "onTokenError: " + e.getLocalizedMessage());
            hmsLogger.sendPeriodicEvent("onTokenError", e.getLocalizedMessage());
        }
    }

    @Override
    public void onNewToken(String token) {
        if (HMSPush.getPlugin() == null)
            return;
        hmsLogger = HMSLogger.getInstance(HMSPush.getPlugin().cordova.getContext());
        try {
            super.onNewToken(token);
            HmsMessagePublisher.sendOnNewTokenEvent(token);
            hmsLogger.sendPeriodicEvent("onNewToken");
        } catch (Exception e) {
            Log.e(TAG, e.getLocalizedMessage());
            hmsLogger.sendPeriodicEvent("onNewToken", e.getLocalizedMessage());
        }
    }
}
