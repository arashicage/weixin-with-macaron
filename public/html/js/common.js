
var menuItem;
function menu_open() {
	menu_close();
	menuItem = $(this).find('ul:first').show();
	if ($(this).index() > 0) {
		$(this).find('a:first').addClass('air');
		if ($(this).find('ul').length > 0) {
			$('.sub_bg').show();
		}
	}
}
function menu_close() {
	if (menuItem) {
		menuItem.hide();
		if ($(this).index() > 0) {
			$(this).find('a:first').removeClass('air');
			$('.sub_bg').hide();
		}
	}
}
$(document).ready(function(){
	$('.nav_main > li').bind('mouseenter', menu_open);
	$('.nav_main > li').bind('mouseleave',  menu_close);
  shownew();
  showTitle();
});

  var NewsTime = 3000;    
  var TextTime = 200;    
  var newTitle =  "支持增值税专用发票、货物运输业增值税专用发票、机动车销售统一发票、增值税普通发票（含电子发票）在线查验"; 

  var newsi = 0;
  var txti = 0;
  var txttimer;   
  var newstimer;
  
  var newstitle = new Array();    
  var newshref = new Array();   
  
  newstitle[0] = newTitle;
  newshref[0] = "";
  
  newstitle[1] = newTitle;
  newshref[1] = "";
  
  newstitle[2] = newTitle;
  newshref[2] = "";
  
  newstitle[3] = newTitle;
  newshref[3] = "";
  
  function shownew(){
    hwnewstr=newstitle[newsi];    
    newslink=newshref[newsi];     
     
    if(txti>=hwnewstr.length){
      clearInterval(txttimer);  
      clearInterval(newstimer); 
      newsi++;   
      
      if(newsi>=newstitle.length){
        newsi = 0; 
      }
      newstimer = setInterval("shownew()",NewsTime);   
      txti = 0;  
      return;
    }
    clearInterval(txttimer);  
    document.getElementById("Hotnews").href = newslink;  
    document.getElementById("Hotnews").innerHTML = hwnewstr.substring(0,txti+1);   
    
    txti++;  
    txttimer = setInterval("shownew()",TextTime);   
  }

  function getGrzx(n){
    var username=getCookie('username');

    var token=getCookie('Qw2#RSDEFRGTGTonj@QSAr3Cd5A6Rg');
    var loginsj=getCookie('dqsj');

    if(username==""||loginsj==""||token==""){
        if(n==0){
          $('#grzx').attr('href','login.html');
        }else{
          $('#grzx').attr('href','../login.html');
        }     
     }else{
        if(n==0){
          $('#grzx').attr('href','xgxx.html');
        }else{
          $('#grzx').attr('href','../xgxx.html');
        }       
     }
  } 

function getzjDz(){
  return "http://172.30.4.36:7001/WebQuery";
}