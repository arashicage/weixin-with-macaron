
function FormatDate(time){
    var year=time.substring(0,4);
    var month=parseInt(time.substring(4,6),10);
    var day=parseInt(time.substring(6),10);
    return year+"年"+month+"月"+day+"日";
}

   function NoToChinese(n,fplx)  
      {
        var flag=0;
        var s = '';  
        var fraction = ['角', '分'];
        var digit = ['零', '壹', '贰', '叁', '肆', '伍', '陆', '柒', '捌', '玖'];
        var unit = [ ['园', '万', '亿'], ['', '拾', '佰', '仟']  ];
        var head = n < 0? '欠': '';
        if(n.split(".")[1]=="00"&&(fplx=="02"||fplx=="03")){
          flag=1;
        }
        if(flag==1){
          var spot=n.split(".")[1];
          s=digit[spot.charAt(0)]+"角"+digit[spot.charAt(1)]+"分";
        }
        
        n = Math.abs(n);          

        if(flag!=1){
          for (var i = 0; i < fraction.length; i++) 
          {
              s += (digit[Math.floor(n * 10 * Math.pow(10, i)) % 10] + fraction[i]).replace(/零./, '');
          }
          s = s || '整';
        } 

        n = Math.floor(n);
    
        for (var i = 0; i < unit[0].length && n > 0; i++) 
        {
            var p = '';
            for (var j = 0; j < unit[1].length && n > 0; j++) 
            {
                p = digit[n % 10] + unit[1][j] + p;
                n = Math.floor(n / 10);
            }
            s = p.replace(/(零.)*零$/, '').replace(/^$/, '零')  + unit[0][i] + s;
        }
        if(flag==1){
          return head + s.replace(/(零.)*零园/, '园');
        }
        return head + s.replace(/(零.)*零园/, '园').replace(/(零.)+/g, '零').replace(/^整$/, '零园整');
  }

  function GetHwxxHtml(hwxxs){
    var hwinfo=hwxxs.split('≡');
    var hw;
    var html = "";

    for(var i=0;i<hwinfo.length;i++){
       
        hw=hwinfo[i].split('█');
        html+='<tr>';
        for(var j=0;j<8;j++){
          if(j!=7){
             html+='<td class="align_center borderRight"><span class="content_td_blue">';
          }else{
            html+='<td class="align_center"><span class="content_td_blue">';
          }          
           if(j==6){
              html+=FormatSl(hw[j]);       
           }else if(j==5){
              html+=GetJeToDot(hw[j]);
           }else if(j==3||j==4||j==7){
            html+=AddDot(hw[j]);
           }
           else{
              html+=hw[j];
           }
           
           html+='</span></td>';
        }
        html+='</tr>';
    }
    return html;
  }

  function GetDzHwxxHtml(hwxxs){
    var hwinfo=hwxxs.split('≡');
    var hw;
    var html = "";

    for(var i=0;i<hwinfo.length;i++){
       
        hw=hwinfo[i].split('█');
        html+='<tr>';
        for(var j=0;j<8;j++){

          if(j!=7){
             html+='<td class="align_center borderRight"><span class="content_td_blue">';
          }else{
            html+='<td class="align_center"><span class="content_td_blue">';
          }  
          if(j==3){
             html+=hw[6];
          }else if(j==6){
              html+=FormatSl(hw[3]);       
          }else{
              html+=hw[j];
          }
           
           html+='</span></td>';
        }
        html+='</tr>';
    }
    return html;
  }

  function FormatSl(data){
    if(data.substring(0,1)=="."){
      data=parseFloat("0"+data)*100;
    }
    return data+"%";
  }

  function AddDot(je){
    var index=je.indexOf(".");
    
    if(index==0){
      je="0"+je;
    }else if(je[0]=="-"&&index==1){
      je="-0."+je.split(".")[1];
    }
    return je;
  }

  function GetJeToDot(je_all){
    var je=AddDot(je_all);
    var index=je.indexOf(".");
    if(index<0){
      je+=".00";
    }else if(je.split(".")[1].length==1){
      je+="0";
    }
    return je;
  }

  function GethyzpHwxxHtml(hwxxs,value){

    var hwinfo=hwxxs.split('≡');
    var hw;
    var html = "";

    for(var i=value;i<hwinfo.length;i=i+2){
        hw=hwinfo[i].split('█');
        html+='<tr>';
        for(var j=0;j<2;j++){
           html+='<td class="align_center"><span class="content_td_blue">';
           if(j==1){
            html+=AddDot(hw[j]);
           }else{
            html+=hw[j];
           }
           
           html+='</span></td>';
        }
        html+='</tr>';
    }
    return html;
  }

  function GetjsfpHwxxHtml(hwxxs,value){
    var hwinfo=hwxxs.split('≡');
    var hw;
    var html = "";

    for(var i=0;i<hwinfo.length;i++){
       
        hw=hwinfo[i].split('█');
        html+='<tr>';
        for(var j=0;j<4;j++){
           html+='<td class="align_center"><span class="content_td_blue">';
           if(j==1||j==2||j==3){
            html+=AddDot(hw[j]);
           }else{
            html+=hw[j];
           }
           
           html+='</span></td>';
        }
        html+='</tr>';
    }
    return html;
  }