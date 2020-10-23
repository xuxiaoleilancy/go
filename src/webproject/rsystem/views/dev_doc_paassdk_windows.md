# 随锐开发文档 Windows

标签（空格分隔）： 随锐 会见通讯云 PaasSDK 20181010

[TOC]

# 1. Windows SDK集成说明

Windows SDK是为在微软Windows操作系统中集成随锐视频功能提供的SDK。使用C++语言开发，用户可以用SDK实现商务沟通功能。目前支持发起会议和加入会议等功能。

## 1.1 集成准备
从[随锐SDK版本管理服务器](http://svn.suirui.com/svn/9-Install/06.SR-huijian/trunk/pc/SRPaasSDK)获取SDK,SDK中目录结构如下:

```
    | -SRPaasSDK
      | - bin
        | - debug
        | - release
      | - include
      | - lib
      | - doc
```
SRPaasSDK目录中包含SDK的头文件、lib文件和dll文件。

## 1.2 集成环境配置
随锐的视频SDK基于Qt5开发。windows下可以选择VS2013作为IDE开发工具，linux下可以直接使用Qt提供的qtcreator。

### 1.2.1 Qt版本说明
随锐SRPaasSDK开发过程中使用的是Qt5.7.0版本发布，当然基于Qt5的任一版本我们都可以提供，如果需要可以发送邮件给我们(developer@suirui.com)。

### 1.2.2 SDK集成
随锐视频SDK以动态库的方式提供，和普通的动态库使用方式一致。

# 2 集成示例
VideoClient是视频服务的入口，可以直接调用VideoClient的接口，也额可以通过VideoClient获得相应的VideoControlManager等进行相应的操作。也可以通过VideoClient注册回调，用于获取接口调用的执行结果。

## 2.1 生成VideoClient实例
```
VideoClient* m_pClient = NULL;

void createVideoClient()
{
	if ( NULL == m_pClient )
	{
	    m_pClient = VideoClient::getInstance();
	    m_pClient->setCallback(this); //用于注册回调
	}
	
}
```
**回调注册[详见2.7](#2.7)**

## 2.2 发起会议

```
//发起会议
void startMeeting()
{
	//appid为第三方集成时向随锐申请的一个产品标识符，可通过web来申请
	QString strAppId = "";
	
	//Uid,用户id，第三方用户唯一标识，数字即可，位数不限
	int Uid = 65789;
	
	/*会议id(如果内网访问不必填,外网需要加@和公司名称)*/
	QString strMeetingNumber = "";
	
	//用户昵称，用于显示在会议中的昵称
	QString strNickName = "pcClient";
	
	//会议密码，如果有会议密码请传入会议密码
	QString strPassword = "";
	
	 //domain:服务器地址(内网IP,外网IP.如果不走外网访问,外网IP为空。例如:https://192.168.61.39/api,https://lab.suirui.com/api)
	QString domain = "";
	
	//serverId暂时不用，可不填（默认值 "123456"）
	QString strServerId = "123456";
	
	//发起会议,内网
	VideoClient::getInstance()->startMeeting(APPID, Uid, "5C3B0DE7A06C596A9EAF3329420E38B5BFFE7736CBF59FC6", strNickName, "", "", "https://192.168.61.39/api");
	//发起会议,外网
	ViedoClient::getInstance()->startMeeting(APPID, Uid, "5C3B0DE7A06C596A9EAF3329420E38B5BFFE7736CBF59FC6", strNickName, "@suirui", "", "https://192.168.61.39/api,https://lab.suirui.com/api");
}
```
## 2.3 加入会议

```
void joinMeeting()
{
	//appid为第三方集成时向随锐申请的一个产品标识符，可通过web来申请
	QString strAppId = "";
	
	//Uid,用户id，第三方用户唯一标识，数字即可，位数不限
	int Uid = 65789;
	
	/*会议号(如果内网访问不必填,外网需要加@和公司名称)*/
	QString strMeetingNumber = "337427";
	
	//用户昵称，用于显示在会议中的昵称
	QString strNickName = "pcClient";
	
	//会议密码，如果有会议密码请传入会议密码
	QString strPassword = "";
	
	//domain:服务器地址(内网IP,外网IP.如果不走外网访问,外网IP为空。例如: https://192.168.61.39/api,https://lab.suirui.com/api)
	QString domain = "";
	
	//serverId暂时不用，可不填（默认值 "123456"）
	QString strServerId = "123456";
	
	//加入会议,内网
	videoClient::getInstance()->joinMeeting("beb7da4ced7c42a085c3c99697f9aa42", Uid, "16ED4496266FD69746076ECF7FDE6E269C991BB7CACA5708", strNickName, strMeetingNumber, "", "https://192.168.61.39/api");
	//加入会议,外网
	videoClient::getInstance()->joinMeeting("beb7da4ced7c42a085c3c99697f9aa42",Uid,"16ED4496266FD69746076ECF7FDE6E269C991BB7CACA5708", strNickName, strMeetingNumber, "", "https://192.168.61.39/api,https://lab.suirui.com/api, "123456", false, true, 1);
}
```
## 2.4 离开会议
```
void leaveMeeting()
{
    //用户离开当前所在会议调用该接口，会弹出提示窗口，请用户确认是否离开会议，点击按钮后实现相应的离会的功能
    g_pClient->leaveMeeting();
}
```

## 2.5 结束会议
```
void endMeeting()
{
    /*如果用户是当前会议的拥有者（即会议主持人），在调用该接口的时候会弹出提示窗口，请用户确认是结束会议还是离开会议，点击按钮后实现相应的离会或结束会议的功能*/
    g_pClient->endMeeting();
}
```
## 2.6 获取当前会议号

```
QString getConferenceID()
{
    //用户成功入会后，调用该接口可以获取当前所在会议的会议号
    QString conferenceID = g_pClient->getConferenceID();
}
```

## 2.7 <span id="2.7"></span>注册回调
注册回调需要实现VideoModuleCallBack接口类的对应接口，并注册到VideoClient中，一般情况下在生成VideoClient的时候注册。

```

class MyVideoCallback : public VideoModuleCallBack
{
public:
	// 发起会议结果回调
	void openMeetingCompleted(bool bSuccessed, const QString &info = QString())
	{
        //bSuccessed  打开会议是否成功，true为打开会议成功，false为打开会议失败
        // info       会议失败的原因
	}
	
	// 离开或结束会议成功的结果回调
	void leaveMeetingComplete(bool isEnd, bool bLeaveBeforeJoinAfter)
	{
        //isEnd  区分当前操作是离开会议还是结束会议（true为结束会议，false为离开会议） 
        //bLeaveBeforeJoinAfter （暂时未用，不需要处理）
	}
	
	// 加入会议时入会连接异常
	void onStackConnectError()
	{
        
	}
	
	//取消离开视频模块的结果回调
	void quitModuleCanceled()
	{

	}
	
	//修改个人会议号是否成功的结果回调（暂时未使用）
	void onChangedOwnConferenceComplete(bool bSuccessed)
	{

	}
}
```

# 3 示例工程
示例工程在公司内网svn服务器，如果需要可向我们免费申请。
[windows SDK集成示例工程](http://svn.suirui.com/svn/9-Install/06.SR-huijian/trunk/pc/SRPaasSDKDemo/)
