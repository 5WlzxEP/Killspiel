import{s as mt,f as k,a as L,g as C,X as M,c as B,h as R,d as x,j as E,i as O,x as i,W as ft,o as gt,l as W,m as Y,n as G,p as lt,e as nt}from"../chunks/scheduler.QGVxIACg.js";import{S as pt,i as bt,a as I,c as Z,t as F,g as tt,b as ut,d as ht,m as dt,e as _t}from"../chunks/index.ElInTfSE.js";import{e as Q}from"../chunks/each.k3i3fjfp.js";import{I as vt,a as kt}from"../chunks/IconX.im0qjcw7.js";import{d as J}from"../chunks/leaderboard_data.JH0A3KKd.js";function st(c,t,s){const e=c.slice();return e[14]=t[s],e}function at(c,t,s){const e=c.slice();return e[17]=t[s],e}function rt(c){let t,s="letzten 8 Versuche";return{c(){t=k("th"),t.textContent=s,this.h()},l(e){t=C(e,"TH",{class:!0,"data-svelte-h":!0}),M(t)!=="svelte-1ey7shy"&&(t.textContent=s),this.h()},h(){E(t,"class","text-right")},m(e,n){O(e,t,n)},d(e){e&&x(t)}}}function ot(c){let t,s,e=Q(c[14].latest.toString(2).padStart(8,"0")),n=[];for(let r=0;r<e.length;r+=1)n[r]=ct(at(c,e,r));const p=r=>F(n[r],1,1,()=>{n[r]=null});return{c(){t=k("td");for(let r=0;r<n.length;r+=1)n[r].c();this.h()},l(r){t=C(r,"TD",{class:!0});var a=R(t);for(let l=0;l<n.length;l+=1)n[l].l(a);a.forEach(x),this.h()},h(){E(t,"class","grid-cols-8 grid")},m(r,a){O(r,t,a);for(let l=0;l<n.length;l+=1)n[l]&&n[l].m(t,null);s=!0},p(r,a){if(a&1){e=Q(r[14].latest.toString(2).padStart(8,"0"));let l;for(l=0;l<e.length;l+=1){const b=at(r,e,l);n[l]?(n[l].p(b,a),I(n[l],1)):(n[l]=ct(b),n[l].c(),I(n[l],1),n[l].m(t,null))}for(tt(),l=e.length;l<n.length;l+=1)p(l);Z()}},i(r){if(!s){for(let a=0;a<e.length;a+=1)I(n[a]);s=!0}},o(r){n=n.filter(Boolean);for(let a=0;a<n.length;a+=1)F(n[a]);s=!1},d(r){r&&x(t),ft(n,r)}}}function Ct(c){let t,s;return t=new vt({props:{color:"lime"}}),{c(){ut(t.$$.fragment)},l(e){ht(t.$$.fragment,e)},m(e,n){dt(t,e,n),s=!0},i(e){s||(I(t.$$.fragment,e),s=!0)},o(e){F(t.$$.fragment,e),s=!1},d(e){_t(t,e)}}}function xt(c){let t,s;return t=new kt({props:{color:"red"}}),{c(){ut(t.$$.fragment)},l(e){ht(t.$$.fragment,e)},m(e,n){dt(t,e,n),s=!0},i(e){s||(I(t.$$.fragment,e),s=!0)},o(e){F(t.$$.fragment,e),s=!1},d(e){_t(t,e)}}}function ct(c){let t,s,e,n;const p=[xt,Ct],r=[];function a(l,b){return l[17]==="0"?0:1}return t=a(c),s=r[t]=p[t](c),{c(){s.c(),e=nt()},l(l){s.l(l),e=nt()},m(l,b){r[t].m(l,b),O(l,e,b),n=!0},p(l,b){let T=t;t=a(l),t!==T&&(tt(),F(r[T],1,1,()=>{r[T]=null}),Z(),s=r[t],s||(s=r[t]=p[t](l),s.c()),I(s,1),s.m(e.parentNode,e))},i(l){n||(I(s),n=!0)},o(l){F(s),n=!1},d(l){l&&x(e),r[t].d(l)}}}function it(c){let t,s,e=c[14].rank+"",n,p,r,a,l=c[14].name+"",b,T,H,A,P=c[14].points+"",y,N,u,_=c[14].guesses+"",S,q,D,V=((c[14].guesses>0?c[14].points/c[14].guesses:0)*100).toFixed(2)+"",X,z,j,K,m,h=c[1]&&ot(c);return{c(){t=k("tr"),s=k("td"),n=W(e),p=L(),r=k("td"),a=k("a"),b=W(l),H=L(),A=k("td"),y=W(P),N=L(),u=k("td"),S=W(_),q=L(),D=k("td"),X=W(V),z=W(" %"),j=L(),h&&h.c(),K=L(),this.h()},l(o){t=C(o,"TR",{});var g=R(t);s=C(g,"TD",{});var f=R(s);n=Y(f,e),f.forEach(x),p=B(g),r=C(g,"TD",{});var v=R(r);a=C(v,"A",{href:!0,class:!0,"data-sveltekit-preload-data":!0});var d=R(a);b=Y(d,l),d.forEach(x),v.forEach(x),H=B(g),A=C(g,"TD",{class:!0});var $=R(A);y=Y($,P),$.forEach(x),N=B(g),u=C(g,"TD",{class:!0});var w=R(u);S=Y(w,_),w.forEach(x),q=B(g),D=C(g,"TD",{class:!0});var U=R(D);X=Y(U,V),z=Y(U," %"),U.forEach(x),j=B(g),h&&h.l(g),K=B(g),g.forEach(x),this.h()},h(){E(a,"href",T="/Killspiel/user/"+c[14].id),E(a,"class","w-full p-4"),E(a,"data-sveltekit-preload-data","tap"),E(A,"class","text-right"),E(u,"class","text-right"),E(D,"class","text-right")},m(o,g){O(o,t,g),i(t,s),i(s,n),i(t,p),i(t,r),i(r,a),i(a,b),i(t,H),i(t,A),i(A,y),i(t,N),i(t,u),i(u,S),i(t,q),i(t,D),i(D,X),i(D,z),i(t,j),h&&h.m(t,null),i(t,K),m=!0},p(o,g){(!m||g&1)&&e!==(e=o[14].rank+"")&&G(n,e),(!m||g&1)&&l!==(l=o[14].name+"")&&G(b,l),(!m||g&1&&T!==(T="/Killspiel/user/"+o[14].id))&&E(a,"href",T),(!m||g&1)&&P!==(P=o[14].points+"")&&G(y,P),(!m||g&1)&&_!==(_=o[14].guesses+"")&&G(S,_),(!m||g&1)&&V!==(V=((o[14].guesses>0?o[14].points/o[14].guesses:0)*100).toFixed(2)+"")&&G(X,V),o[1]?h?(h.p(o,g),g&2&&I(h,1)):(h=ot(o),h.c(),I(h,1),h.m(t,K)):h&&(tt(),F(h,1,1,()=>{h=null}),Z())},i(o){m||(I(h),m=!0)},o(o){F(h),m=!1},d(o){o&&x(t),h&&h.d()}}}function Tt(c){let t,s="Leaderboard",e,n,p,r,a,l,b="Rank",T,H,A="Name",P,y,N="Punkte",u,_,S="Teilnahmen",q,D,V="Rate",X,z,j,K,m=c[1]&&rt(),h=Q(c[0]),o=[];for(let f=0;f<h.length;f+=1)o[f]=it(st(c,h,f));const g=f=>F(o[f],1,1,()=>{o[f]=null});return{c(){t=k("h1"),t.textContent=s,e=L(),n=k("div"),p=k("table"),r=k("thead"),a=k("tr"),l=k("th"),l.textContent=b,T=L(),H=k("th"),H.textContent=A,P=L(),y=k("th"),y.textContent=N,u=L(),_=k("th"),_.textContent=S,q=L(),D=k("th"),D.textContent=V,X=L(),m&&m.c(),z=L(),j=k("tbody");for(let f=0;f<o.length;f+=1)o[f].c();this.h()},l(f){t=C(f,"H1",{class:!0,"data-svelte-h":!0}),M(t)!=="svelte-nu7ra6"&&(t.textContent=s),e=B(f),n=C(f,"DIV",{class:!0});var v=R(n);p=C(v,"TABLE",{class:!0});var d=R(p);r=C(d,"THEAD",{});var $=R(r);a=C($,"TR",{});var w=R(a);l=C(w,"TH",{"data-svelte-h":!0}),M(l)!=="svelte-1ccjcls"&&(l.textContent=b),T=B(w),H=C(w,"TH",{class:!0,"data-svelte-h":!0}),M(H)!=="svelte-duhd4o"&&(H.textContent=A),P=B(w),y=C(w,"TH",{class:!0,"data-svelte-h":!0}),M(y)!=="svelte-byy2g0"&&(y.textContent=N),u=B(w),_=C(w,"TH",{class:!0,"data-svelte-h":!0}),M(_)!=="svelte-1t5aoyg"&&(_.textContent=S),q=B(w),D=C(w,"TH",{class:!0,"data-svelte-h":!0}),M(D)!=="svelte-17o6624"&&(D.textContent=V),X=B(w),m&&m.l(w),w.forEach(x),$.forEach(x),z=B(d),j=C(d,"TBODY",{});var U=R(j);for(let et=0;et<o.length;et+=1)o[et].l(U);U.forEach(x),d.forEach(x),v.forEach(x),this.h()},h(){E(t,"class","text-primary-500 text-center text-lg card rounded-none"),E(H,"class","text-center"),E(y,"class","text-center"),E(_,"class","text-center"),E(D,"class","text-center"),E(p,"class","table table-hover table-cell-fit mx-auto w-full rounded-none"),E(n,"class","table-container mx-auto w-full rounded-none")},m(f,v){O(f,t,v),O(f,e,v),O(f,n,v),i(n,p),i(p,r),i(r,a),i(a,l),i(a,T),i(a,H),c[5](H),i(a,P),i(a,y),c[6](y),i(a,u),i(a,_),c[7](_),i(a,q),i(a,D),i(a,X),m&&m.m(a,null),i(p,z),i(p,j);for(let d=0;d<o.length;d+=1)o[d]&&o[d].m(j,null);K=!0},p(f,[v]){if(f[1]?m||(m=rt(),m.c(),m.m(a,null)):m&&(m.d(1),m=null),v&3){h=Q(f[0]);let d;for(d=0;d<h.length;d+=1){const $=st(f,h,d);o[d]?(o[d].p($,v),I(o[d],1)):(o[d]=it($),o[d].c(),I(o[d],1),o[d].m(j,null))}for(tt(),d=h.length;d<o.length;d+=1)g(d);Z()}},i(f){if(!K){for(let v=0;v<h.length;v+=1)I(o[v]);K=!0}},o(f){o=o.filter(Boolean);for(let v=0;v<o.length;v+=1)F(o[v]);K=!1},d(f){f&&(x(t),x(e),x(n)),c[5](null),c[6](null),c[7](null),m&&m.d(),ft(o,f)}}}function Et(c,t,s){let e=J,n="p",p="asc",r=!1;async function a(){const u=p=="asc"?1:-1;switch(n){case"n":s(0,e=J.sort((_,S)=>u*(_.name<S.name?-1:_.name==S.name?0:1)));break;case"p":s(0,e=J.sort((_,S)=>u*(_.points-S.points)));break;case"g":s(0,e=J.sort((_,S)=>u*(_.guesses-S.guesses)));break}}let l,b,T;gt(()=>{H(),a()});function H(){const u=new URLSearchParams(location.search);p=u.get("order")==="desc"?"desc":"asc",A(u.get("sorting")),s(1,r=u.get("latest")==="true"),u.get("limit")}function A(u){let _;switch(u){case"name":n="n",_=l;break;case"teilnahmen":n="g",_=T;break;default:n="p",_=b}p==="asc"?_.classList.add("table-sort-asc"):_.classList.add("table-sort-dsc")}function P(u){lt[u?"unshift":"push"](()=>{l=u,s(2,l)})}function y(u){lt[u?"unshift":"push"](()=>{b=u,s(3,b)})}function N(u){lt[u?"unshift":"push"](()=>{T=u,s(4,T)})}return[e,r,l,b,T,P,y,N]}class Lt extends pt{constructor(t){super(),bt(this,t,Et,Tt,mt,{})}}export{Lt as component};
