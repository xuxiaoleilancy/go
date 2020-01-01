var debugTextArea = document.getElementById("debugTextArea");
var displayRSelectVideoBusiness = document.getElementById("displayRSelectVideoBusiness");

 function format(txt,compress/*是否为压缩模式*/){/* 格式化JSON源码(对象转换为JSON文本) */  
        var indentChar = '    ';   
        if(/^\s*$/.test(txt)){   
            alert('数据为空,无法格式化! ');   
            return;   
        }   
        try{var data=eval('('+txt+')');}   
        catch(e){   
            alert('数据源语法错误,格式化失败! 错误信息: '+e.description,'err');   
            return;   
        };   
        var draw=[],last=false,This=this,line=compress?'':'\n',nodeCount=0,maxDepth=0;   
           
        var notify=function(name,value,isLast,indent/*缩进*/,formObj){   
            nodeCount++;/*节点计数*/  
            for (var i=0,tab='';i<indent;i++ )tab+=indentChar;/* 缩进HTML */  
            tab=compress?'':tab;/*压缩模式忽略缩进*/  
            maxDepth=++indent;/*缩进递增并记录*/  
            if(value&&value.constructor==Array){/*处理数组*/  
                draw.push(tab+(formObj?('"'+name+'":'):'')+'['+line);/*缩进'[' 然后换行*/  
                for (var i=0;i<value.length;i++)   
                    notify(i,value[i],i==value.length-1,indent,false);   
                draw.push(tab+']'+(isLast?line:(','+line)));/*缩进']'换行,若非尾元素则添加逗号*/  
            }else   if(value&&typeof value=='object'){/*处理对象*/  
                    draw.push(tab+(formObj?('"'+name+'":'):'')+'{'+line);/*缩进'{' 然后换行*/  
                    var len=0,i=0;   
                    for(var key in value)len++;   
                    for(var key in value)notify(key,value[key],++i==len,indent,true);   
                    draw.push(tab+'}'+(isLast?line:(','+line)));/*缩进'}'换行,若非尾元素则添加逗号*/  
                }else{   
                        if(typeof value=='string')value='"'+value+'"';   
                        draw.push(tab+(formObj?('"'+name+'":'):'')+value+(isLast?'':',')+line);   
                };   
        };   
        var isLast=true,indent=0;   
        notify('',data,isLast,indent,false);   
        return draw.join('');   
    }

function formatJson(json, options) 
{
    var reg = null,
            formatted = '',
            pad = 0,
            PADDING = '    ';
    options = options || {};
    options.newlineAfterColonIfBeforeBraceOrBracket = (options.newlineAfterColonIfBeforeBraceOrBracket === true) ? true : false;
    options.spaceAfterColon = (options.spaceAfterColon === false) ? false : true;
    if (typeof json !== 'string') {
        json = JSON.stringify(json);
    } else {
        json = JSON.parse(json);
        json = JSON.stringify(json);
    }
    reg = /([\{\}])/g;
    json = json.replace(reg, '\r\n$1\r\n');
    reg = /([\[\]])/g;
    json = json.replace(reg, '\r\n$1\r\n');
    reg = /(\,)/g;
    json = json.replace(reg, '$1\r\n');
    reg = /(\r\n\r\n)/g;
    json = json.replace(reg, '\r\n');
    reg = /\r\n\,/g;
    json = json.replace(reg, ',');
    if (!options.newlineAfterColonIfBeforeBraceOrBracket) {
        reg = /\:\r\n\{/g;
        json = json.replace(reg, ':{');
        reg = /\:\r\n\[/g;
        json = json.replace(reg, ':[');
    }
    if (options.spaceAfterColon) {
        reg = /\:/g;
        json = json.replace(reg, ':');
    }
    (json.split('\r\n')).forEach(function (node, index) {
                var i = 0,
                        indent = 0,
                        padding = '';

                if (node.match(/\{$/) || node.match(/\[$/)) {
                    indent = 1;
                } else if (node.match(/\}/) || node.match(/\]/)) {
                    if (pad !== 0) {
                        pad -= 1;
                    }
                } else {
                    indent = 0;
                }

                for (i = 0; i < pad; i++) {
                    padding += PADDING;
                }

                formatted += padding + node + '\r\n';
                pad += indent;
            }
    );
    return formatted;
}

function showJsonFormat(msg){
    //var resultJson = formatJson(msg);
    //document.getElementById("writePlace").innerHTML = '<pre>' +resultJson + '<pre/>';
    var y = formatJson(msg,true);  
    $("#jsonformatresult").empty();  
    $("#jsonformatresult").append("<textarea id='jsondata' class='jsondatatextarea'>" + y + "</textarea>");  
    //$('#result').html(syntaxHighlight( format(msg,true) ) );
}
function debug(message) {
    //result = JSON.stringify(message, null, 2);
    //debugTextArea.value += result + "\n";
    //var json = $('#result').val();
    //var result = JSON.stringify(JSON.parse(json), null, 4);
    //$('#result').val(syntaxHighlight(result));

    //
    //(3)将格式化好后的json写入页面中
    //document.getElementById("debugTextArea").value = resultJson ;
    if (!displayRSelectVideoBusiness.checked) {
        var strusiness = "RSelectVideoBusiness"
        if (message.indexOf(strusiness) != -1) {
            return
        }
    }

    debugTextArea.value += "--------------------------------------------------------------------------------------\n" + "-----" + message + "\n" + "-----------------------------------------------------------------------------------------------------\n"
    //$('#result').html(syntaxHighlight(message));
    debugTextArea.scrollTop = debugTextArea.scrollHeight;
}

function clearMsg() {
    debugTextArea.value = "";
}
var orgURL=window.location.host;
var hostIP=orgURL.substring(0,orgURL.indexOf(':'));
debug(hostIP);

/*
function getEndpointData(endpoint){
    var send_cmd='{"header":{"id":0,"sign":"'+endpoint+'","endpoint":"'+endpoint+'","srcendpoint":"'+endpoint+'","from":"websocket",\
        "traces":[],\
        "schema":0,\
        "action":0\
        },\
        "data":""}'
    if (websocket != null) {
         websocket.send( send_cmd );
    }
    console.log(send_cmd); 
}
*/
function getIntStringData(ivalue,sValue){
    return '"{\\\"ivalue\\\":'+ivalue+',\\\"svalue\\\":\\\"'+sValue+'\\\"}"'
}
function getIntData(ivalue){
    return '"{\\\"value\\\":'+ivalue+'\\\"}"'
}
function getEndpointData(endpoint,data){

    var send_cmd='{"header":{"id":0,"sign":"'+endpoint+'","endpoint":"'+endpoint+'","to":"rtcvideo","from":"client",\
        "traces":[],\
        "schema":0,\
        "action":0\
        },\
        "data":'+data+'}'
    if (websocket != null) {
         websocket.send( send_cmd );
    }
    console.log(send_cmd); 
}
function getEndpointDataFromPaasUI(endpoint,data){

    var send_cmd='{"header":{"id":0,"sign":"'+endpoint+'","endpoint":"'+endpoint+'","to":"paasui"'+',"from":"client",\
        "traces":[],\
        "schema":0,\
        "action":0\
        },\
        "data":'+data+'}'
    if (websocket != null) {
         websocket.send( send_cmd );
    }
    console.log(send_cmd); 
}
function updateEndpointData(endpoint,data){

    var send_cmd='{"header":{"id":0,"sign":"'+endpoint+'","endpoint":"'+endpoint+'","srcendpoint":"'+endpoint+'","from":"client",\
        "traces":[],\
        "schema":0,\
        "action":1\
        },\
        "result":{"code":0,"des":"rinvaliddes"},\
        "data":'+data+'}'
    if (websocket != null) {
         websocket.send( send_cmd );
    }
    console.log(send_cmd); 
}
function getterdata(endpoint){

    var terid =document.getElementById("terinput").value;

    if ( terid == 0 ) {
        return '""'
    }else{
         return getIntStringData(terid,'')
    }

    return '""'
}

function renderchannels(){
    var endpoint = 'terminals:id'
    getEndpointDataFromPaasUI(endpoint,'""')
}
function getmodevalue(){
    var mode =document.getElementById("modeinput").value;
    var mode_type = 0
    if ( mode == 'mode1' ) {
        mode_type = 10000 + 1;
    }

    if ( mode == 'mode2' ) {
        mode_type = 10000 + 2;
    }
    
    if ( mode == 'mode3' ) {
        mode_type = 10000 + 3;
    }
    return mode_type
}
function getmodedata(endpoint){

    var mode_type = getmodevalue()
   
    if ( mode_type == 0 ) {
        return '""'
    }

    if ( endpoint == 'terminals') {
        return '""'
    }

    if ( endpoint == 'rect2terms') {
        return getIntStringData(mode_type,'')
    }
    if ( endpoint == "layouts:") {
         return getIntStringData(mode_type,'')
    }
    if ( endpoint == "conferences:id:vtable:realselects") {
         return '""'
    }
    if ( endpoint == "conferences:id:vtable:modeselects") {
         return getIntData(mode_type)
    }
    return '""'
}

function selecteds() {
    var endpoint = 'conferences:id:vtable:realselects'
    getEndpointData(endpoint,'""')
}

function modeselect() {
    var endpoint = 'conferences:id:vtable:modeselects'
    var data = getIntData( getmodevalue())
    getEndpointData(endpoint,data)
}
function terminals() {
    var endpoint = 'terminals'
    var data = getterdata(endpoint)
    if ( data != '""') {
        endpoint = 'terminals:id'
    }
    getEndpointData(endpoint,data)
}

function showLocalPreview() {
    var endpoint = 'test:localvideopreview';
	var data = '""';
	getEndpointDataFromPaasUI(endpoint,data);
}
function getscreens() {
    var endpoint = 'action:device:screeninfo';
    var data = '""';
    getEndpointDataFromPaasUI(endpoint, data);
}

function rectterminals( ) {
    var endpoint = 'rect2terms'
    var data = getmodedata(endpoint)
    if ( data != '""') {
        endpoint = 'rect2terms:id'
    }
    getEndpointData(endpoint,data)
}
function moderects() {

    var endpoint = 'layouts:'
    var data = getmodedata(endpoint)
    if ( data != '""') {
        endpoint = 'layouts:id'
    }
    getEndpointData(endpoint,data)
}
function conferenceIndex() {
    var endpoint = 'conferences:id:index'
    getEndpointData(endpoint,'""')
}
function conferenceState() {

    var endpoint = 'conferences:id:states'
    getEndpointData(endpoint,'""')
}
function conferenceActivity() {

    var endpoint = 'layouts:'
    getEndpointData(endpoint,'""')
}

function setlayoutbyflag(layoutflag){

	var data = '"{\\\"displayid\\\":1,\\\"rectList\\\":[],\\\"layouttype\\\":' + layoutflag + ',\\\"toSet\\\":\\\"1\\\"}"';
    //var data = '"{\\\"mode\\\":'+10001+',\\\"layout\\\":'+layoutflag+'}"'
    var endpoint = 'layouts:id:operation'
    updateEndpointData(endpoint,data)
}

function setlayout(){
    var debugTextArea = document.getElementById("layoutcountinput");
    var count = debugTextArea.value

    setlayoutbyflag(count)

}

function getcurrentlayout(type){
		var data = '"{\\\"value\\\":\\\"' + type + '\\\"}"';
        var endpoint = 'layouts:id:current'
		getEndpointData(endpoint,data)
}

function setlayout_auto(){
    var flag = 0
    setlayoutbyflag(flag)
}
function setlayout_lecture(){
    var flag = 105
    setlayoutbyflag(flag)
}
function changconfmode(type){
	var data = '"{\\\"value\\\":' + type +  '}"';
    var endpoint = 'conferences:id:states:confmode';
    updateEndpointData(endpoint,data)
}
function setMasterlayout(){
	var debugTextArea = document.getElementById("layoutcountinput");
    var count = debugTextArea.value
	var data = '"{\\\"layouttype\\\":' + count + ',\\\"toSet\\\":\\\"2\\\"}"';
	var endpoint = 'layouts:id:operation';
	updateEndpointData(endpoint,data);
}
function sendMessage() {
    var msg = document.getElementById("inputText").value;
     if ( websocket != null )
    {
        websocket.send(msg);
        console.log( startcmd);
    }
    
}

var wsUri;
var websocket = null;

function initWebSocket() {
    try {
        if (typeof MozWebSocket == 'function')
            WebSocket = MozWebSocket;
        if ( websocket && websocket.readyState == 1 )
            websocket.close();
        var serverip=document.getElementById("ip").value;
        if (serverip=='') {
            debug("don't have url");
            wsUri = "ws://localhost:9898";
            debug(wsUri);
        } else {
             wsUri = 'ws://'+serverip+':9898';
             debug(wsUri);
        }
        websocket = new WebSocket( wsUri );
        websocket.onopen = function (evt) {
            debug("CONNECTED");
            showSuccessNotify("COnnected");
        };
        websocket.onclose = function (evt) {
            debug("DISCONNECTED");
        };
        websocket.onmessage = function (evt) {
            // console.log( "Message received :", evt.data );
            debug( evt.data );
            showJsonFormat(evt.data)
        };
        websocket.onerror = function (evt) {
            debug('ERROR: ' + evt.data);
            debug('can not connect to Server:');
        };
    } catch (exception) {
        debug('ERROR: ' + exception);
    }
}
function handleWSError(ev) {
    debug('ERROR: ' +  msg);
}

function stopWebSocket() {
    if (websocket)
        websocket.close();
}

function checkSocket() {
    if (websocket != null) {
        var stateStr;
        switch (websocket.readyState) {
            case 0: {
                stateStr = "CONNECTING";
                break;
            }
            case 1: {
                stateStr = "OPEN";
                debug(window.navigator.userAgent);
                debug("test");
                break;
            }
            case 2: {
                stateStr = "CLOSING";
                break;
            }
            case 3: {
                stateStr = "CLOSED";
                break;
            }
            default: {
                stateStr = "UNKNOW";
                break;
            }
        }
        debug("WebSocket state = " + websocket.readyState + " ( " + stateStr + " )");
    } else {
        debug("WebSocket is null");
    }


}

$(function () {

       // $("#mySelect").select(); 不传参数可以这样写
        $("#mySelect").select({
            ontSize:"185px", 
            width: "200px",

        });
        
        //可选参数,不填就是默认值
        //width: "180px",            //生成的select框宽度
        //listMaxHeight:"200px",     //生成的下拉列表最大高度
        //themeColor: "#00bb9c",    //主题颜色
        //fontColor: "#000",        //字体颜色
        //fontFamily: "'Helvetica Neue', arial, sans-serif",    //字体种类
        //fontSize:"15px",           //字体大小
        //showSearch: false,        //是否启用搜索框
        //rowColor:"#fff",          //行原本的颜色
        //rowHoverColor: "#0faf03", //移动选择时，每一行的hover底色
        //fontHoverColor: "#fff",   //移动选择时，每一行的字体hover颜色
        //mainContent: "请选择",    //选择显示框的默认文字
        //searchContent: "关键词搜索"   //搜索框的默认提示文字  
    });
