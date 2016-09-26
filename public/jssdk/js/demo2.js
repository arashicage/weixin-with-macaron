/*
 * 注意：
 * 1. 所有的JS接口只能在公众号绑定的域名下调用，公众号开发者需要先登录微信公众平台进入“公众号设置”的“功能设置”里填写“JS接口安全域名”。
 * 2. 如果发现在 Android 不能分享自定义内容，请到官网下载最新的包覆盖安装，Android 自定义分享接口需升级至 6.0.2.58 版本及以上。
 * 3. 完整 JS-SDK 文档地址：http://mp.weixin.qq.com/wiki/7/aaa137b55fb2e0456bf8dd9148dd613f.html
 *
 * 如有问题请通过以下渠道反馈：
 * 邮箱地址：weixin-open@qq.com
 * 邮件主题：【微信JS-SDK反馈】具体问题
 * 邮件内容说明：用简明的语言描述问题所在，并交代清楚遇到该问题的场景，可附上截屏图片，微信团队会尽快处理你的反馈。
 */
wx.ready(function () {

    // 隐藏右上角菜单
    wx.hideOptionMenu();

    document.querySelector('#scanQRCode').onclick = function () {
        wx.scanQRCode({
            needResult: 1,
            desc: '扫描和识别发票的二维码',
            scanType: ["qrCode"],
            success: function (res) {

                var fields = res.resultStr.split(',');

                if (fields.length != 7) {
                    $('#errMsg').html("非法或无法识别的二维码，请重新扫描。");
                    $('#dialog2').show();
                    $('#dialog2').find('.weui_btn_dialog').on('click', function () {
                        $('#dialog2').hide();
                    });
                    return
                } else if (fields[1].length != 2) { //检查 发票类型
                    $('#errMsg').html("无效的发票类型，请重新扫描。");
                    $('#dialog2').show();
                    $('#dialog2').find('.weui_btn_dialog').on('click', function () {
                        $('#dialog2').hide();
                    });
                    return
                } else if (fields[2].length != 10 && fields[2].length != 12) { //检查 发票代码
                    document.getElementById("kpdm").value = fields[2];
                    $('#errMsg').html("无效的发票代码，请重新扫描。");
                    $('#dialog2').show();
                    $('#dialog2').find('.weui_btn_dialog').on('click', function () {
                        $('#dialog2').hide();
                    });
                    return
                } else if (fields[3].length != 8) { //检查 发票号码
                    $('#errMsg').html("无效的发票号码，请重新扫描。");
                    $('#dialog2').show();
                    $('#dialog2').find('.weui_btn_dialog').on('click', function () {
                        $('#dialog2').hide();
                    });
                    return
                }/*else if (fields[4].length != 2){ //检查 开票金额
                 $('#errMsg').html("无效的开票金额，请重新扫描。");
                 $('#dialog2').show();
                 $('#dialog2').find('.weui_btn_dialog').on('click', function () {
                 $('#dialog2').hide();
                 });
                 return
                 }else if (fields[5].length != 2){ //检查 开票日期
                 $('#errMsg').html("无效的开票日期，请重新扫描。");
                 $('#dialog2').show();
                 $('#dialog2').find('.weui_btn_dialog').on('click', function () {
                 $('#dialog2').hide();
                 });
                 return
                 }*/ else {
                    //document.getElementById("fplx").value = fields[1];
                    document.getElementById("fpdm").value = fields[2];
                    document.getElementById("fphm").value = fields[3];
                    document.getElementById("kpje").value = fields[4];
                    document.getElementById("kprq").value = fields[5].substr(0, 4) + "-" + fields[5].substr(4, 2) + "-" + fields[5].substr(6, 2);
                }

                //alert(JSON.stringify(res));
            }
        });
    };

});

wx.error(function (res) {

    alert(res.errMsg);
});

