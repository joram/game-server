(this.webpackJsonpgame=this.webpackJsonpgame||[]).push([[0],{33:function(e,t,a){e.exports=a(48)},38:function(e,t,a){},39:function(e,t,a){},48:function(e,t,a){"use strict";a.r(t);var n=a(0),o=a.n(n),s=a(26),i=a.n(s),r=(a(38),a(3)),c=a(2),l=a(5),u=a(4),p=a(27),h=a.n(p),d=(a(39),a(28)),m=a.n(d),v=a(15),f=a.n(v),y=a(20),b=a.n(y),g=a(29);function k(){var e=window.location.href.split("/")[2];return"localhost:3000"===e?"localhost:2305":e}function E(){return"localhost:2305"===k()?"ws":"wss"}function j(){return"localhost:2305"===k()?"http":"https"}function x(e){return new Promise((function(t){return setTimeout(t,e)}))}var O=new(function(){function e(){var t=this;Object(r.a)(this,e),this.objectEventBus=a(45)();var n="".concat(E(),"://").concat(k(),"/objects");console.log(n),this.ws=new WebSocket(n),this.ws.onopen=function(){console.log("connected to generic ws")},this.ws.onmessage=function(e){var a=JSON.parse(e.data);t.processMessage(a)},this.ws.onclose=function(){console.log("disconnected generic ws")}}return Object(c.a)(e,[{key:"sendAccessToken",value:function(e,t,a,n,o){var s=this;this.waitForConn((function(){var i='{"accessToken":"'.concat(e,'", "googleId":"').concat(t,'", "email":"').concat(a,'", "firstName":"').concat(n,'", "lastName":"').concat(o,'"}');s.ws.send(i)}))}},{key:"sendBackpackRequest",value:function(){var e=this;this.waitForConn((function(){e.ws.send('{"backpack":"please"}')}))}},{key:"sendPlayerMove",value:function(e){var t=this;this.waitForConn((function(){t.ws.send('{"direction":"'.concat(e,'"}'))}))}},{key:"equipItem",value:function(e){var t=this;this.waitForConn((function(){t.ws.send('{"equip_item":'.concat(e,"}"))}))}},{key:"unequipItem",value:function(e){var t=this;this.waitForConn((function(){t.ws.send('{"unequip_item":'.concat(e,"}"))}))}},{key:"dropItem",value:function(e){var t=this;this.waitForConn((function(){t.ws.send('{"drop_item":'.concat(e,"}"))}))}},{key:"waitForConn",value:function(){var e=Object(g.a)(b.a.mark((function e(t){return b.a.wrap((function(e){for(;;)switch(e.prev=e.next){case 0:if(1===this.ws.readyState){e.next=5;break}return e.next=3,x(1);case 3:e.next=0;break;case 5:t();case 6:case"end":return e.stop()}}),e,this)})));return function(t){return e.apply(this,arguments)}}()},{key:"processMessage",value:function(e){void 0!==e.equipped_image?this.objectEventBus.emit("item",null,e):void 0!==e.playerId?this.objectEventBus.emit("player_id",null,e):void 0!==e.type?this.objectEventBus.emit("monster",null,e):console.log("unknown msg",e)}}]),e}()),w=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){var e;Object(r.a)(this,a);for(var n=arguments.length,o=new Array(n),s=0;s<n;s++)o[s]=arguments[s];return(e=t.call.apply(t,[this].concat(o))).state={id:void 0,playerPosition:{x:0,y:0}},e.debouncedMovePlayer=m.a.debounce((function(t,a,n){e.movePlayer(t,a)}),350,{leading:!0,maxWait:350}),e}return Object(c.a)(a,[{key:"componentDidMount",value:function(){var e=this;O.objectEventBus.on("player_id",(function(t){var a=e.state;a.id=t.playerId,e.setState(a)})),O.objectEventBus.on("monster",(function(t){if(t.id===e.state.id){var a=e.state;a.playerPosition.x=t.x,a.playerPosition.y=t.y,e.setState(a)}var n=e.state,o=n.playerPosition;if("player"===t.type){n.playerPosition=o,n.nextPlayerPosition=o,e.setState(n),e.props.background.current.updateChunks();var s=e.props.app.state;s.playerPosition=o,e.props.app.setState(s)}}))}},{key:"movePlayer",value:function(e,t){"e"===e?this.props.app.setPage("inventory"):("w"===e&&(e="up"),"a"===e&&(e="left"),"s"===e&&(e="down"),"d"===e&&(e="right"),this.props.background.current.updateChunks(),O.sendPlayerMove(e))}},{key:"render",value:function(){var e=this;return o.a.createElement(o.a.Fragment,null,o.a.createElement(f.a,{handleKeys:["left","right","up","down","w","a","s","d","e"],onKeyEvent:function(t,a){e.debouncedMovePlayer(t,a,100)}}),this.props.children)}}]),a}(o.a.Component),S=a(23),C=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){return Object(r.a)(this,a),t.apply(this,arguments)}return Object(c.a)(a,[{key:"isSolid",value:function(){return this.props.pixel.g>180}},{key:"render",value:function(){var e=this.props.pixel,t="rgb(".concat(e.r,", ").concat(e.g,", ").concat(e.b,")");return this.isSolid()&&(t="rgb(0, 80, 0)"),o.a.createElement("div",{style:{width:"50px",height:"50px",left:"".concat(50*(e.x-this.props.chunkPosition.x),"px"),top:"".concat(50*(e.y-this.props.chunkPosition.y),"px"),backgroundColor:t,position:"absolute"}})}}]),a}(o.a.Component),I=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){var e;Object(r.a)(this,a);for(var n=arguments.length,o=new Array(n),s=0;s<n;s++)o[s]=arguments[s];return(e=t.call.apply(t,[this].concat(o))).SIZE=50,e.chunk={},e.state={},e}return Object(c.a)(a,[{key:"backgroundSquares",value:function(){var e=this,t=[];return this.props.data.pixels.forEach((function(a){t.push(o.a.createElement(C,{key:a.x+"_"+a.y,pixel:a,chunkPosition:{x:e.props.data.x,y:e.props.data.y},playerPosition:e.props.playerPosition}))})),t}},{key:"render",value:function(){var e=this.props.x,t=this.props.y,a=Math.ceil(window.innerHeight/50),n=Math.ceil(window.innerWidth/50),s=-this.props.playerPosition.x+Math.floor(n/2),i=-this.props.playerPosition.y+Math.floor(a/2);return o.a.createElement("div",{className:"chunk",style:{position:"absolute",border:"solid thin black",width:"".concat(this.SIZE*this.chunk.size,"px"),height:"".concat(this.SIZE*this.chunk.size,"px"),backgroundColor:"pink",left:50*(e+s)+"px",top:50*(t+i)+"px"}},o.a.createElement("div",{className:"squares"},this.backgroundSquares()))}}]),a}(o.a.Component),N=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){var e;Object(r.a)(this,a);for(var n=arguments.length,o=new Array(n),s=0;s<n;s++)o[s]=arguments[s];return(e=t.call.apply(t,[this].concat(o))).CHUNKSIZE=10,e.state={playerPosition:{x:-99,y:-99},chunkData:{}},e}return Object(c.a)(a,[{key:"componentDidMount",value:function(e){this.updateChunks()}},{key:"requestChunk",value:function(e,t){var a=this;if(void 0===this.state.chunkData["".concat(e,"_").concat(t)]){var n=this.state;n.chunkData["".concat(e,"_").concat(t)]={pixels:[],objects:[]},this.setState(n);var o="".concat(j(),"://").concat(k(),"/chunks?x=").concat(e,"&y=").concat(t);fetch(o).then((function(e){return e.json()})).then((function(n){var o=a.state;o.chunkData["".concat(e,"_").concat(t)]=n,a.setState(o)}))}}},{key:"updateChunks",value:function(){var e=this,t=Math.ceil(window.innerHeight/50),a=Math.ceil(window.innerWidth/50),n=this.playerPosition(),o=Math.floor(n.x-a/2),s=Math.floor(n.y-t/2),i=Math.ceil(n.x+a/2),r=Math.ceil(n.y+t/2);o=o-o%this.CHUNKSIZE-this.CHUNKSIZE,s=s-s%this.CHUNKSIZE-this.CHUNKSIZE,i=i+i%this.CHUNKSIZE+this.CHUNKSIZE,r=r+r%this.CHUNKSIZE+this.CHUNKSIZE;var c=Math.ceil((i-o)/this.CHUNKSIZE),l=Math.ceil((r-s)/this.CHUNKSIZE),u=Object(S.a)(Array(c).keys()),p=Object(S.a)(Array(l).keys());u.forEach((function(t){p.forEach((function(a){void 0===e.state.chunkData["".concat(t,"_").concat(a)]&&e.requestChunk(o+t*e.CHUNKSIZE,s+a*e.CHUNKSIZE)}))}))}},{key:"playerPosition",value:function(){return this.props.app.playerPosition()}},{key:"isSolidAt",value:function(e){var t=Math.floor(e.x/this.CHUNKSIZE)*this.CHUNKSIZE,a=Math.floor(e.y/this.CHUNKSIZE)*this.CHUNKSIZE,n=this.state.chunkData["".concat(t,"_").concat(a)];if(void 0===n)return!1;var o=!1;return n.pixels.forEach((function(t){Math.ceil(e.x)===t.x&&Math.ceil(e.y)===t.y&&t.g>180&&(o=!0)})),o}},{key:"render",value:function(){var e=this,t=[];return Object.keys(this.state.chunkData).forEach((function(a){var n=e.state.chunkData[a],s=o.a.createElement(I,{x:n.x,y:n.y,key:a,playerPosition:e.playerPosition(),data:n,objectEventBus:e.props.objectEventBus});t.push(s)})),o.a.createElement("div",{id:"chunks"},t)}}]),a}(o.a.Component),M=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(e){var n;return Object(r.a)(this,a),(n=t.call(this,e)).state={monster:{x:0,y:0}},n.state.monster=e.monster,n}return Object(c.a)(a,[{key:"isSolid",value:function(){return this.props.monster.solid}}]),Object(c.a)(a,[{key:"componentDidMount",value:function(){var e=this;O.objectEventBus.on("monster",(function(t){if(e.props.id===t.id){var a=e.state;a.monster=t,e.setState(a)}}))}},{key:"render",value:function(){var e=this,t=this.state.monster,a=this.props.size,n=t.x,s=t.y,i=Math.ceil(window.innerHeight/50),r=Math.ceil(window.innerWidth/50),c=-this.props.playerPosition.x+Math.floor(r/2),l=-this.props.playerPosition.y+Math.floor(i/2),u=[];return t.images.forEach((function(n){u.push(o.a.createElement("img",{style:{width:"".concat(a,"px"),height:"".concat(a,"px"),position:"absolute",left:"0px",top:"0px"},key:"".concat(t.id,"_img_").concat(u.length),src:"".concat(j(),"://").concat(k()).concat(n),alt:e.props.monster.type}))})),o.a.createElement("div",{style:{position:"absolute",width:a+"px",height:a+"px",zIndex:this.props.zIndex,left:a*(n+c)+"px",top:a*(s+l)+"px"}},u)}}]),a}(o.a.Component),P=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){var e;Object(r.a)(this,a);for(var n=arguments.length,o=new Array(n),s=0;s<n;s++)o[s]=arguments[s];return(e=t.call.apply(t,[this].concat(o))).SIZE=50,e.state={monsters:{}},e}return Object(c.a)(a,[{key:"componentDidMount",value:function(){var e=this;O.objectEventBus.on("monster",(function(t){var a=e.state;"remove"===t.action?delete a.monsters[t.id]:a.monsters[t.id]=t,e.setState(a)}))}},{key:"isSolidAt",value:function(e){var t=!1;return Object.values(this.state.monsters).forEach((function(a){Math.ceil(e.x)===a.x&&Math.ceil(e.y)===a.y&&(t=t||a.solid)})),console.log(t),t}},{key:"render",value:function(){var e=this,t={x:0,y:0};function a(e,a,n){return o.a.createElement(M,{monster:e,key:e.id,id:e.id,size:a,objectEventBus:n,playerPosition:t})}null!==this.props.player.current&&(t=this.props.player.current.state.playerPosition);var n=[];return Object.values(this.state.monsters).forEach((function(t){"player"!==t.type&&n.push(a(t,e.SIZE,e.props.objectEventBus))})),Object.values(this.state.monsters).forEach((function(t){"player"===t.type&&n.push(a(t,e.SIZE,e.props.objectEventBus))})),o.a.createElement("div",{className:"objects"},n)}}]),a}(o.a.Component),_=a(47).v4,D=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){return Object(r.a)(this,a),t.apply(this,arguments)}return Object(c.a)(a,[{key:"render",value:function(){return o.a.createElement("div",{className:"chat_message",style:{color:"white"}},this.props.message.from.substr(1,6),": ",this.props.message.message)}}]),a}(o.a.Component),q=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){var e;Object(r.a)(this,a);for(var n=arguments.length,o=new Array(n),s=0;s<n;s++)o[s]=arguments[s];return(e=t.call.apply(t,[this].concat(o))).id=_(),e.ws=new WebSocket("".concat(E(),"://").concat(k(),"/chat")),e.state={messages:[],text:""},e}return Object(c.a)(a,[{key:"componentDidMount",value:function(){var e=this;this.ws.onopen=function(){e.ws.send('{"id":"'.concat(e.id,'", "message":"logged in"}'))},this.ws.onmessage=function(t){var a=JSON.parse(t.data),n=e.state;n.messages.push(a),e.setState(n)},this.ws.onclose=function(){console.log("disconnected")}}},{key:"sendMessage",value:function(e,t){e.preventDefault(),console.log(this.state.text),this.ws.send('{"id":"'.concat(this.id,'", "message":"').concat(this.state.text,'"}'));var a=this.state;a.text="",this.setState(a)}},{key:"textChange",value:function(e){var t=this.state;t.text=e.target.value,this.setState(t)}},{key:"render",value:function(){var e=[];return this.state.messages.forEach((function(t,a){e.push(o.a.createElement(D,{key:"msg_".concat(a),message:t}))})),o.a.createElement("div",{style:{width:"300px",height:"500px",backgroundColor:"rgba(20,20,20,0.8)",position:"absolute",top:"10px",left:"10px"}},o.a.createElement("div",{style:{color:"white"}},"you are ",this.id.substr(1,6)),o.a.createElement("div",null,e),o.a.createElement("form",{onSubmit:this.sendMessage.bind(this)},o.a.createElement("input",{type:"text",style:{bottom:"6px",left:"5px",width:"185px",position:"absolute"},onChange:this.textChange.bind(this),value:this.state.text})))}}]),a}(o.a.Component),K=a(9),Z=a(13),B=a(50),H=a(51),U=a(31),A=a(53),F=a(52),T={1:"s1",2:"s2",3:"s3",4:"s4",5:"s5",6:"s6",7:"s7",8:"s8",BACKPACK:"backpack",GROUND:"ground"},z=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(e){var n;return Object(r.a)(this,a),n=t.call(this,e),O.sendBackpackRequest(),n}return Object(c.a)(a,[{key:"render",value:function(){var e=this;return o.a.createElement(o.a.Fragment,null,o.a.createElement("div",{style:{textAlign:"center"}},o.a.createElement("div",{style:{float:"left",width:"300px"}},"Equipped"),o.a.createElement("div",{style:{float:"left",width:"300px"}},"Backpack"),o.a.createElement("div",{style:{float:"left",width:"300px"}},"Ground")),o.a.createElement("div",{style:{top:"20px",position:"absolute"}},o.a.createElement(B.a,{backend:U.a},o.a.createElement(L,null),o.a.createElement(X,null),o.a.createElement($,null))),o.a.createElement(f.a,{handleKeys:["esc","e"],onKeyEvent:function(t,a){e.props.app.setPage("game")}}))}}]),a}(o.a.Component);function W(e){var t=Object(A.a)({item:{type:T[e.item.allowed_slot]},end:function(t,a){var n=a.getDropResult();t&&n&&("backpack"===n.name?O.unequipItem(e.item.id):"ground"===n.name?O.dropItem(e.item.id):O.equipItem(e.item.id))},collect:function(e){return{isDragging:e.isDragging()}}}),a=Object(Z.a)(t,3),n=(a[0].isDragging,a[1]),s=a[2];return o.a.createElement(o.a.Fragment,null,o.a.createElement(H.a,{connect:s,src:e.item.dropped_image}),o.a.createElement("div",{ref:n,style:Object(K.a)({},e.style)},o.a.createElement("img",{src:"".concat(j(),"://").concat(k(),"/").concat(e.item.dropped_image),alt:e.item.name})))}function R(e){var t=Object(F.a)({accept:T[e.slot],drop:function(){return{name:"slot ".concat(e.slot)}},collect:function(e){return{isOver:e.isOver(),canDrop:e.canDrop()}}}),a=Object(Z.a)(t,2),n=a[0],s=n.canDrop,i=n.isOver,r=a[1],c="#222";s&&i?c="darkgreen":s&&(c="darkkhaki");var l=null;return void 0!==e.item&&(l=o.a.createElement(W,{item:e.item,style:Object(K.a)(Object(K.a)({},e.style),{},{top:"10px",left:"10px"})})),o.a.createElement(o.a.Fragment,null,o.a.createElement("div",{ref:r,style:Object(K.a)(Object(K.a)({},e.style),{},{backgroundColor:c})},o.a.createElement("img",{style:{width:"50px",height:"50px"},src:"".concat(j(),"://").concat(k(),"/images/gui/tab_unselected.png"),alt:"slot ".concat(e.slot)})),l)}function G(e){var t=Object(F.a)({accept:Object.values(T),drop:function(){return{name:"backpack"}},collect:function(e){return{isOver:e.isOver(),canDrop:e.canDrop()}}}),a=Object(Z.a)(t,2),n=a[0],s=n.canDrop,i=n.isOver,r=a[1],c="#222";s&&i?c="darkgreen":s&&(c="darkkhaki");return o.a.createElement(o.a.Fragment,null,o.a.createElement("div",{ref:r,style:{width:"100%",height:"100%",border:"solid thin black",backgroundColor:c}}),null)}function J(e){var t=Object(F.a)({accept:Object.values(T),drop:function(){return{name:"ground"}},collect:function(e){return{isOver:e.isOver(),canDrop:e.canDrop()}}}),a=Object(Z.a)(t,2),n=a[0],s=n.canDrop,i=n.isOver,r=a[1],c="#222";s&&i?c="darkgreen":s&&(c="darkkhaki");return o.a.createElement(o.a.Fragment,null,o.a.createElement("div",{ref:r,style:{width:"100%",height:"100%",border:"solid thin black",backgroundColor:c}}),null)}var X=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){var e;Object(r.a)(this,a);for(var n=arguments.length,o=new Array(n),s=0;s<n;s++)o[s]=arguments[s];return(e=t.call.apply(t,[this].concat(o))).state={items:{}},e}return Object(c.a)(a,[{key:"componentDidMount",value:function(){var e=this;O.objectEventBus.on("item",(function(t){var a=e.state;t.is_carried&&!t.is_equipped?a.items[t.id]=t:delete a.items[t.id],e.setState(a)}))}},{key:"render",value:function(){var e=[];return Object.values(this.state.items).forEach((function(t){e.push(o.a.createElement(W,{item:t,key:t.id}))})),o.a.createElement("div",{style:{left:"300px",width:"300px",height:"290px",position:"absolute",border:"solid thin black"}},o.a.createElement(G,null),o.a.createElement("div",{style:{position:"absolute",top:0,left:0}},e))}}]),a}(o.a.Component),L=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){var e;Object(r.a)(this,a);for(var n=arguments.length,o=new Array(n),s=0;s<n;s++)o[s]=arguments[s];return(e=t.call.apply(t,[this].concat(o))).state={items:{}},e}return Object(c.a)(a,[{key:"componentDidMount",value:function(){var e=this;O.objectEventBus.on("item",(function(t){var a=e.state;t.is_carried&&t.is_equipped?a.items[t.equipped_slot]=t:delete a.items[t.allowed_slot],e.setState(a)}))}},{key:"render",value:function(){var e={position:"absolute",width:"200px",height:"200px",left:"50px",top:"40px"},t=[];return Object.values(this.state.items).forEach((function(a){t.push(o.a.createElement("img",{style:e,src:"".concat(j(),"://").concat(k(),"/").concat(a.equipped_image),alt:"silhouette",key:a.id}))})),o.a.createElement("div",{style:{width:"300px",height:"290px",position:"absolute",border:"solid thin black"}},o.a.createElement("img",{style:e,src:"".concat(j(),"://").concat(k(),"/images/player/base/human_m.png"),alt:"silhouette"}),t,o.a.createElement(R,{style:{position:"absolute",left:"0px"},slot:1,item:this.state.items[1]}),o.a.createElement(R,{style:{position:"absolute",transform:"scaleX(-1)",right:"0px"},slot:2,item:this.state.items[2]}),o.a.createElement("div",{style:{position:"absolute",top:"80px",width:"100%"}},o.a.createElement(R,{style:{position:"absolute",left:"0px"},slot:3,item:this.state.items[3]}),o.a.createElement(R,{style:{position:"absolute",transform:"scaleX(-1)",right:"0px"},slot:4,item:this.state.items[4]})),o.a.createElement("div",{style:{position:"absolute",top:"160px",width:"100%"}},o.a.createElement(R,{style:{position:"absolute",left:"0px"},slot:5}),o.a.createElement(R,{style:{position:"absolute",transform:"scaleX(-1)",right:"0px"},slot:6})),o.a.createElement("div",{style:{position:"absolute",top:"240px",width:"100%"}},o.a.createElement(R,{style:{position:"absolute",left:"0px"},slot:7}),o.a.createElement(R,{style:{position:"absolute",transform:"scaleX(-1)",right:"0px"},slot:8})))}}]),a}(o.a.Component),$=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(){var e;Object(r.a)(this,a);for(var n=arguments.length,o=new Array(n),s=0;s<n;s++)o[s]=arguments[s];return(e=t.call.apply(t,[this].concat(o))).state={items:{}},e}return Object(c.a)(a,[{key:"componentDidMount",value:function(){var e=this;O.objectEventBus.on("item",(function(t){var a=e.state;t.is_carried?(delete a.items[t.allowed_slot],delete a.items[t.equipped_slot],delete a.items[-1]):a.items[t.equipped_slot]=t,e.setState(a)}))}},{key:"render",value:function(){var e=[];return Object.values(this.state.items).forEach((function(t){e.push(o.a.createElement(W,{item:t,key:t.id}))})),o.a.createElement("div",{style:{left:"600px",width:"300px",height:"290px",position:"absolute",border:"solid thin black"}},o.a.createElement(J,null),o.a.createElement("div",{style:{position:"absolute",top:0,left:0}},e))}}]),a}(o.a.Component),Q=z,V=function(e){Object(l.a)(a,e);var t=Object(u.a)(a);function a(e){var n;return Object(r.a)(this,a),(n=t.call(this,e)).state={accessToken:null,googleId:null,email:"",firstName:"",lastName:"",page:"game"},n.background=o.a.createRef(),n.objects=o.a.createRef(),n.player=o.a.createRef(),n}return Object(c.a)(a,[{key:"playerPosition",value:function(){return null===this.player.current?{x:0,y:0}:this.player.current.state.playerPosition}},{key:"loginSuccess",value:function(e){var t=this.state;t.accessToken=e.accessToken,t.googleId=e.googleId,t.firstName=e.profileObj.givenName,t.lastName=e.profileObj.familyName,t.email=e.profileObj.email,O.sendAccessToken(t.accessToken,t.googleId,t.email,t.firstName,t.lastName),this.setState(t)}},{key:"renderLogin",value:function(){return o.a.createElement("div",{style:{margin:"auto",width:"200px",paddingTop:"200px"}},o.a.createElement(h.a,{clientId:"662193159992-4fv4hq3q25mkerlt0eqqr1ii670ogugr.apps.googleusercontent.com",onSuccess:this.loginSuccess.bind(this),isSignedIn:!0}))}},{key:"setPage",value:function(e){var t=this.state;t.page=e,console.log(t),this.setState(t)}},{key:"renderGame",value:function(){return o.a.createElement(o.a.Fragment,null,o.a.createElement(w,{size:50,background:this.background,objects:this.objects,ref:this.player,app:this,accessToken:this.state.accessToken,googleId:this.state.googleId,email:this.state.email,firstName:this.state.firstName,lastName:this.state.lastName}),o.a.createElement(N,{size:50,ref:this.background,app:this}),o.a.createElement(P,{size:50,player:this.player,ref:this.objects}),o.a.createElement(q,null))}},{key:"renderInventory",value:function(){return o.a.createElement(Q,{app:this,player:this.player,accessToken:this.state.accessToken,googleId:this.state.googleId,email:this.state.email,firstName:this.state.firstName,lastName:this.state.lastName})}},{key:"render",value:function(){return null===this.state.accessToken?this.renderLogin():"game"===this.state.page?this.renderGame():"inventory"===this.state.page?this.renderInventory():null}}]),a}(o.a.Component);Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));i.a.render(o.a.createElement(o.a.StrictMode,null,o.a.createElement(V,null)),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(e){e.unregister()})).catch((function(e){console.error(e.message)}))}},[[33,1,2]]]);
//# sourceMappingURL=main.189fe49e.chunk.js.map