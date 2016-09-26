var adm=function(a){
	var b=/^1\d{11}$|^\d{6}[1-9]\d{2}0$/;
	var c=/^0{8}[1-9]?\w[1-9]\d*$/;

	var e=b.test(a);
	var f=c.test(a);

	if(e==true&&bc(a)&&alxd(a)!="99"){
		return true;
	}else{
		return false;
	}
};

var bc=function(a){
	var b;
	var d=new Date();
	var e=d.getFullYear();
	var f=e.toString();
	var g=f.substring(2);
	if(a.length==12){
		b=a.substring(5,7);
	}else{
		b=a.substring(4,6);
	}

	if(b<=00||b>g){
		return false;
	}
	return true;
};

var ahm=function(a){
	var b=/^\d{8}$/;
	var c=/^0{8}$/;
	var d=/^0{11}-?[1-9]*\w\d*$/;
	var e=b.test(a);
	var f=c.test(a);
	if(e==true&&f==false){
		return true;
	}else{
		return false;
	}
};

var acq=function(a){

	var b=/^\d{8}$/;
	var c=/^0{8}$/;
	var d=/^0{11}-?[1-9]*\w\d*$/;
	var e=b.test(a);
	if(e==true){
		var g=new Date();
		var h=g.getFullYear();
		var i=g.getMonth()+1;
		var j=g.getDate();
		var k=a.substring(0,4);
		var l=a.substring(4,6);
		var m=a.substring(6);
		var n=ca(0);
		var t=ca(1);
		if((h!=k&&h-1!=k)||l>12||m>31||a>n||!cb(h,l,m)){
			return false;
		}		
		if(h-1==k&&a<=t){
			return false;
		}
		return true;
	}else{
		return false;
	}
};


function ca(i){
	var a = new Date();
	var b = 0;
	var c = 0;
	var d = 0;
	var e = "";
	b= a.getFullYear()-i;
	c= a.getMonth()+1;
	d = a.getDate();
	e += b;
	if (c >= 10 ){
		e += c;
	}else{
		e += "0" + c;
	}
	if (d >= 10 ){
		e += d ;
	}else{
		e += "0" + d ;
	}
	return e;
}

function cb(a,b,c){
	if(b==2){
		if(c>29){
			return false;
		}
		if (((a % 4)==0) && ((a % 100)!=0) || ((a % 400)==0)) {
			if(c>28){
				return false;
			}
		}
	}else if((b==4||b==6||b==9||b==11)&&c>30){
		return false;
	}
	return true;
}

var alxd=function(a){
	var b;
	var c="99";
	var code=new Array('144031539110','131001570151','133011501118','111001571071');

	if(a.length==12){
		b=a.substring(7,8);

		for(var i=0;i<code.length;i++){
			if(a==code[i]){
				c="10";
				break;
			}
		}
		if(c=="99"){
			if(b==2){
				c="03";
			}else{
				c="11";
			}
		}		 

	}else if(a.length==10){
		b=a.substring(7,8);
		if(b==1||b==5){
			c="01";
		}else if(b==6||b==3){
			c="04";
		}else if(b==7||b==2){
			c="02";
		}
	}
	return c;
};

var aje=function(a,b){
	var c=alxd(a);
	if(c=="01"||c=="02"||c=="03"){
		return ea(b);
	}else{
		return eb(b);
	}
};

var ea=function(a){
	var b=/(^-?\d{1,11}$)|(^-?\d{1,11}\.\d{1,2}$)/;
	var c=b.test(a);
	return c;
};

var eb=function(a){
	var b=/^-?(\d+$)|(\d+\.\d{1,2})$/;
	var c=b.test(a);
	return c;
};

var ajy=function(a){
	var b=/^\d{6}$/;
	var e=b.test(a);
	return e;
};

function avym(yzm){
  var code;
  for(var i=0;i<yzm.length;i++){
    code=yzm.charCodeAt(i);
    if((code>65248)||(code==12288)){
      return false;
    }
  }
  return true;
}

function aept(fpdm,fphm,kprq,kjje,yzm){

    if(fpdm==""||fphm==""||kprq==""||kjje==""||yzm==""){      
       return false;
    }
    return true;        
}

function arw(){
 	$('#fpdm').val("");
    $('#fphm').val("");
    $('#kprq').val("YYYYMMDD");
    $('#kprq').css('color','#999999');
    $('#kjje').val("");
    $('#yzm').val("请输入验证码或答案");
    $('#yzm').css('color','#999999');
}

function avai(fplx,xsje){

	var fpdm=$("#fpdm").val().trim();
	var fphm=$("#fphm").val().trim();
	var kprq=$("#kprq").val().trim();
	var kjje=$("#kjje").val().trim(); 
	var jym=$("#jym").val().trim();
	var yzm=$("#yzm").val().trim();

	if(kprq=="YYMMDD"){
		kprq="";
	}
	if(yzm=="请输入验证码或答案"){
		yzm="";
	}

	if(xsje==1&&!aept(fpdm,fphm,kprq,kjje,yzm)){
		return false;
	}

	if(xsje==0&&!aept(fpdm,fphm,kprq,jym,yzm)){
		return false;
	}
	
	if(fplx=="99"){
		return false;
	}

    if(adm(fpdm)&&ahm(fphm)&&acq(kprq)&&((xsje==1&&aje(fpdm,kjje))||(xsje==0&&ajy(jym)))&&avym(yzm)){
    	return true;
    }
    return false;

}

function acb(fplx,xsje){

    if(avai(fplx,xsje)){
      $('#uncheckfp').hide();
      $('#checkfp').show();
    }else{
       $('#checkfp').hide();
      $('#uncheckfp').show();
    }
}


function getSwjg(fpdm){
  var citys=[{'code':'1100','sfmc':'北京'},
             {'code':'1200','sfmc':'天津'},
             {'code':'1300','sfmc':'河北'},
             {'code':'1400','sfmc':'山西'},
             {'code':'1500','sfmc':'内蒙古'},
             {'code':'2100','sfmc':'辽宁'},
             {'code':'2102','sfmc':'大连'},
             {'code':'2200','sfmc':'吉林'},
             {'code':'2300','sfmc':'黑龙江'},
             {'code':'3100','sfmc':'上海'},
             {'code':'3200','sfmc':'江苏'},
             {'code':'3300','sfmc':'浙江'},
             {'code':'3302','sfmc':'宁波'},
             {'code':'3400','sfmc':'安徽'},
             {'code':'3500','sfmc':'福建'},
             {'code':'3502','sfmc':'厦门'},
             {'code':'3600','sfmc':'江西'},
             {'code':'3700','sfmc':'山东'},
             {'code':'3702','sfmc':'青岛'},
             {'code':'4100','sfmc':'河南'},
             {'code':'4200','sfmc':'湖北'},
             {'code':'4300','sfmc':'湖南'},
             {'code':'4400','sfmc':'广东'},
             {'code':'4403','sfmc':'深圳'},
             {'code':'4500','sfmc':'广西'},
             {'code':'4600','sfmc':'海南'},
             {'code':'5000','sfmc':'重庆'},
             {'code':'5100','sfmc':'四川'},
             {'code':'5200','sfmc':'贵州'},
             {'code':'5300','sfmc':'云南'},
             {'code':'5400','sfmc':'西藏'},
             {'code':'6100','sfmc':'陕西'},
             {'code':'6200','sfmc':'甘肃'},
             {'code':'6300','sfmc':'青海'},
             {'code':'6400','sfmc':'宁夏'},
             {'code':'6500','sfmc':'新疆'}];
  var dqdm=null;
  var swjginfo=new Array();

  if(fpdm.length==12){
    dqdm=fpdm.substring(1,5);
  }else{
    dqdm=fpdm.substring(0,4);
  }
  if(dqdm!="2102"&&dqdm!="3302"&&dqdm!="3502"&&dqdm!="3702"&&dqdm!="4403"){
    dqdm=dqdm.substring(0,2)+"00";
  }
  for(var i=0;i<citys.length;i++){
    if(dqdm==citys[i].code){
      swjginfo[0]=citys[i].sfmc;
     //swjginfo[1]="http://130.9.1.47:7001/WebQuery";
     	  swjginfo[1]="http://187837.cicp.net:23183/WebQuery/do/"+dqdm;
    //  swjginfo[1]="https://fpcy.gszj.gov.cn/WebQuery/do/"+dqdm;
      break;
    }
  }
  return swjginfo;
}
