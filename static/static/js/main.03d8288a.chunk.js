(this.webpackJsonpgame=this.webpackJsonpgame||[]).push([[0],{14:function(t,e,o){},15:function(t,e,o){},21:function(t,e,o){"use strict";o.r(e);var a=o(0),n=o.n(a),i=o(6),r=o.n(i),s=(o(14),o(1)),l=o(2),c=o(4),p=o(3),h=(o(15),function(t){Object(c.a)(o,t);var e=Object(p.a)(o);function o(){return Object(s.a)(this,o),e.apply(this,arguments)}return Object(l.a)(o,[{key:"render",value:function(){var t=this.props.x,e=this.props.y,o=Math.ceil(window.innerHeight/50),a=Math.ceil(window.innerWidth/50),i=-this.props.playerPosition.x+Math.floor(a/2),r=-this.props.playerPosition.y+Math.floor(o/2);return n.a.createElement("div",{style:{position:"absolute",width:"50px",height:"50px",zIndex:this.props.zIndex,left:50*(t+i)+"px",top:50*(e+r)+"px"}},this.props.children)}}]),o}(n.a.Component)),u=function(t){Object(c.a)(o,t);var e=Object(p.a)(o);function o(){return Object(s.a)(this,o),e.apply(this,arguments)}return Object(l.a)(o,[{key:"isSolid",value:function(){return this.props.pixel.g>180}},{key:"render",value:function(){var t=this.props.pixel,e="rgb(".concat(t.r,", ").concat(t.g,", ").concat(t.b,")");return this.isSolid()&&(e="rgb(0, 80, 0)"),n.a.createElement(h,{x:t.x,y:t.y,playerPosition:this.props.playerPosition},n.a.createElement("div",{style:{width:"100%",height:"100%",backgroundColor:e}}))}}]),o}(n.a.Component),y=function(t){Object(c.a)(o,t);var e=Object(p.a)(o);function o(){var t;Object(s.a)(this,o);for(var a=arguments.length,n=new Array(a),i=0;i<a;i++)n[i]=arguments[i];return(t=e.call.apply(e,[this].concat(n))).state={pixels:[],squares:[],playerPosition:{x:-99,y:-99}},t}return Object(l.a)(o,[{key:"updateSquares",value:function(){var t=this;if(Math.floor(this.props.playerPosition.x)!==Math.floor(this.state.playerPosition.x)||Math.floor(this.props.playerPosition.y)!==Math.floor(this.state.playerPosition.y)){var e=Math.ceil(window.innerHeight/50),o=Math.ceil(window.innerWidth/50),a=Math.floor(this.props.playerPosition.x)-Math.floor(o/2),n=Math.floor(this.props.playerPosition.x)+Math.ceil(o/2)+2,i=Math.floor(this.props.playerPosition.y)-Math.floor(e/2),r=Math.floor(this.props.playerPosition.y)+Math.ceil(e/2)+2,s="http://localhost:2303/pixels?x1=".concat(a-1,"&y1=").concat(i,"&x2=").concat(n,"&y2=").concat(r+1);console.log(s),fetch(s).then((function(t){return t.json()})).then((function(e){var o=t.state;o.pixels=e,o.playerPosition.x=t.props.playerPosition.x,o.playerPosition.y=t.props.playerPosition.y,t.setState(o)}))}}},{key:"squares",value:function(){var t=this,e=[];return this.state.pixels.forEach((function(o){e.push(n.a.createElement(u,{key:o.x+"_"+o.y,pixel:o,playerPosition:t.props.playerPosition}))})),e}},{key:"isSolidAt",value:function(t){var e=!1;return this.state.pixels.forEach((function(o){o.x===Math.ceil(t.x)&&o.y===Math.ceil(t.y)&&(e=o.g>180)})),e}},{key:"render",value:function(){return this.updateSquares(),n.a.createElement("div",{style:{overflow:"hidden",position:"absolute",backgroundColor:"red",width:"100%",height:"100%"}},this.squares())}}]),o}(n.a.Component),f=o(7),d=o.n(f),v=o(8),x=o.n(v);function b(t){var e=t.size;return n.a.createElement(h,{x:t.x,y:t.y,playerPosition:t.playerPosition,zIndex:10},n.a.createElement("div",{style:{backgroundColor:"black",borderRadius:"50%",width:e+"px",height:e+"px"}}))}var P=function(t){Object(c.a)(o,t);var e=Object(p.a)(o);function o(){var t;Object(s.a)(this,o);for(var a=arguments.length,n=new Array(a),i=0;i<a;i++)n[i]=arguments[i];return(t=e.call.apply(e,[this].concat(n))).state={playerPosition:{x:0,y:0},nextPlayerPosition:{x:0,y:0}},t.debouncedMovePlayer=d.a.debounce((function(e,o,a){t.movePlayer(e,o)}),150,{leading:!0,maxWait:150}),t}return Object(l.a)(o,[{key:"movePlayer",value:function(t,e){var o={x:this.state.playerPosition.x,y:this.state.playerPosition.y};function a(t,e,o){return["left","a"].includes(e)&&(t.x-=o),"right"!==e&&"d"!==e||(t.x+=o),"up"!==e&&"w"!==e||(t.y-=o),"down"!==e&&"s"!==e||(t.y+=o),t}["left","a","up","w"].includes(t)&&(o=a(o,t,1)),o=a(o,t,.1);var n=this.props.squares.current.isSolidAt(o);o=a(o,t,.1),["left","a","up","w"].includes(t)&&(o=a(o,t,-1));var i=Math.round(o.x/.2),r=Math.round(o.y/.2);if(o.x=.2*i,o.y=.2*r,console.log(o),!n){var s=this.state;s.playerPosition=o,s.nextPlayerPosition=o,this.setState(s);var l=this.props.app.state;l.playerPosition=o,this.props.app.setState(l),this.props.app.setPlayerPosition(o)}}},{key:"render",value:function(){var t=this;return n.a.createElement(n.a.Fragment,null,n.a.createElement(b,{x:this.state.playerPosition.x,y:this.state.playerPosition.y,playerPosition:this.state.playerPosition,size:this.props.size}),n.a.createElement(x.a,{handleKeys:["left","right","up","down","w","a","s","d"],onKeyEvent:function(e,o){t.debouncedMovePlayer(e,o,100)}}),this.props.children)}}]),o}(n.a.Component),m=function(t){Object(c.a)(o,t);var e=Object(p.a)(o);function o(){return Object(s.a)(this,o),e.apply(this,arguments)}return Object(l.a)(o,[{key:"isSolid",value:function(){return this.props.object.solid}},{key:"render",value:function(){var t=this.props.object,e=this.props.size,o=t.x,a=t.y;return n.a.createElement("div",{style:{position:"absolute",width:e+"px",height:e+"px",zIndex:this.props.zIndex,left:e*o+this.props.offsetX+"px",top:e*a+this.props.offsetY+"px"}},n.a.createElement("img",{src:this.props.object.image,alt:this.props.object.type}))}}]),o}(n.a.Component),j=function(t){Object(c.a)(o,t);var e=Object(p.a)(o);function o(){var t;Object(s.a)(this,o);for(var a=arguments.length,n=new Array(a),i=0;i<a;i++)n[i]=arguments[i];return(t=e.call.apply(e,[this].concat(n))).state={},t.hasGotObjects=!1,t}return Object(l.a)(o,[{key:"getObjects",value:function(){var t=this;if(!this.hasGotObjects){var e="http://localhost:2303/chunks?x=".concat(this.props.x,"&y=").concat(this.props.y);fetch(e).then((function(t){return t.json()})).then((function(e){var o=t.state;o.chunk=e,t.setState(o),t.hasGotObjects=!0}))}}},{key:"render",value:function(){if(this.getObjects(),!this.hasGotObjects)return n.a.createElement("div",null);var t=Math.ceil(window.innerHeight/50),e=Math.ceil(window.innerWidth/50),o=this.props.playerPosition.x-Math.floor(e/2),a=this.props.playerPosition.y-Math.floor(t/2),i=[],r=this.state.chunk.objects;return r!=={}&&void 0!==r||(r=[]),r.forEach((function(t){i.push(n.a.createElement(m,{object:t,key:t.id,size:50,offsetX:50*-o,offsetY:50*(-a-1)}))})),n.a.createElement("div",{style:{zIndex:"1",position:"absolute"}},i)}}]),o}(n.a.Component),g=function(t){Object(c.a)(o,t);var e=Object(p.a)(o);function o(t){var a;return Object(s.a)(this,o),(a=e.call(this,t)).state={},a.squares=n.a.createRef(),a.player=n.a.createRef(),a}return Object(l.a)(o,[{key:"playerPosition",value:function(){return null===this.player.current?{x:0,y:0}:this.player.current.state.playerPosition}},{key:"setPlayerPosition",value:function(t){console.log("setting player pos",t)}},{key:"render",value:function(){return n.a.createElement("div",null,n.a.createElement(P,{size:50,squares:this.squares,ref:this.player,app:this}),n.a.createElement(j,{x:0,y:0,playerPosition:this.playerPosition()}),n.a.createElement(y,{playerPosition:this.playerPosition(),ref:this.squares}))}}]),o}(n.a.Component);Boolean("localhost"===window.location.hostname||"[::1]"===window.location.hostname||window.location.hostname.match(/^127(?:\.(?:25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}$/));r.a.render(n.a.createElement(n.a.StrictMode,null,n.a.createElement(g,null)),document.getElementById("root")),"serviceWorker"in navigator&&navigator.serviceWorker.ready.then((function(t){t.unregister()})).catch((function(t){console.error(t.message)}))},9:function(t,e,o){t.exports=o(21)}},[[9,1,2]]]);
//# sourceMappingURL=main.03d8288a.chunk.js.map