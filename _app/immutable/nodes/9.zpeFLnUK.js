import{s as R,a as D,f as b,S as N,d as h,c as z,g as y,h as k,j,i as O,x as d,y as A,X as I,l as x,m as K,n as L,W,F as X}from"../chunks/scheduler.QGVxIACg.js";import{e as F}from"../chunks/each.k3i3fjfp.js";import{S as Y,i as J}from"../chunks/index.ElInTfSE.js";import{g as Q}from"../chunks/navigation.fExWMQKf.js";const U=[{name:"tammie dennis dds_3183",id:73999208,vote:10},{name:"kurt everett dvm_5602",id:36314529,vote:10},{name:"nicole kerr_7952",id:25362043,vote:10},{name:"carla gillespie dds_8518",id:10342805,vote:10},{name:"dr. chase blair_6350",id:76309099,vote:10},{name:"dr. jerome doyle_202",id:87901346,vote:10},{name:"dr. dale gates_4305",id:23642742,vote:10},{name:"dr. kari cortez dds_5347",id:99068210,vote:10},{name:"melody pacheco_6548",id:43418300,vote:10},{name:"dr. jay trevino_8645",id:5287066,vote:10},{name:"mr. travis morgan_546",id:74757793,vote:10},{name:"dr. rose fitzgerald md_1355",id:30008971,vote:10},{name:"maurice warner_4727",id:69678003,vote:10},{name:"dr. cristian choi_7920",id:35707954,vote:10},{name:"wesley bonilla v_3365",id:64210792,vote:10},{name:"miss lacey carroll_1017",id:9901934,vote:10},{name:"ashley stout_5381",id:54272856,vote:10},{name:"miss mandy rowland md_8972",id:96196105,vote:10},{name:"miss krista ortega_5674",id:49361368,vote:10},{name:"richard oneill dvm_4517",id:29388465,vote:10},{name:"dean daniel phd_3453",id:81543885,vote:10},{name:"mr. parker benson jr._9396",id:99935782,vote:10},{name:"dominique powell_5179",id:70180782,vote:10},{name:"ms. selena warren_7382",id:59242308,vote:10},{name:"mark prince ii_4223",id:6936651,vote:10},{name:"ruth wolfe md_7450",id:24713344,vote:10},{name:"marissa montoya_1702",id:22147144,vote:10},{name:"mrs. valerie woods md_514",id:16671289,vote:10},{name:"joanna atkinson dvm_1330",id:55934939,vote:10},{name:"mr. kent acevedo_7715",id:36816622,vote:10},{name:"ryan mooney ii_1759",id:55878886,vote:10},{name:"chelsea hunter_6801",id:16753884,vote:10},{name:"dr. aaron gilmore dds_3095",id:94718345,vote:10},{name:"mrs. savannah raymond_3569",id:63453205,vote:10},{name:"mrs. judith orozco_3225",id:15036903,vote:10},{name:"morgan leon phd_5922",id:21262249,vote:10},{name:"brent pham_5394",id:36690356,vote:10},{name:"mr. willie underwood md_8115",id:47458973,vote:10}],Z=!1,$=async({params:l})=>({data:U,params:l}),se=Object.freeze(Object.defineProperty({__proto__:null,load:$,prerender:Z},Symbol.toStringTag,{value:"Module"}));function G(l,e,a){const t=l.slice();return t[2]=e[a],t}function ee(l){let e,a="Keine Ergebnisse gefunden";return{c(){e=b("div"),e.textContent=a,this.h()},l(t){e=y(t,"DIV",{class:!0,"data-svelte-h":!0}),I(e)!=="svelte-vaw0mv"&&(e.textContent=a),this.h()},h(){j(e,"class","p-4")},m(t,r){O(t,e,r)},p:A,d(t){t&&h(e)}}}function te(l){var M;let e,a,t,r='<tr><th>Name</th> <th class="text-right">Vote</th></tr>',m,c,v,o,n,i,T="Total",w,p,f=((M=l[0].data)==null?void 0:M.length)+"",C,E=F(l[0].data),_=[];for(let s=0;s<E.length;s+=1)_[s]=P(G(l,E,s));return{c(){e=b("div"),a=b("table"),t=b("thead"),t.innerHTML=r,m=D(),c=b("tbody");for(let s=0;s<_.length;s+=1)_[s].c();v=D(),o=b("tfoot"),n=b("tr"),i=b("th"),i.textContent=T,w=D(),p=b("th"),C=x(f),this.h()},l(s){e=y(s,"DIV",{class:!0});var S=k(e);a=y(S,"TABLE",{class:!0});var g=k(a);t=y(g,"THEAD",{"data-svelte-h":!0}),I(t)!=="svelte-134eqbq"&&(t.innerHTML=r),m=z(g),c=y(g,"TBODY",{class:!0});var u=k(c);for(let q=0;q<_.length;q+=1)_[q].l(u);u.forEach(h),v=z(g),o=y(g,"TFOOT",{});var H=k(o);n=y(H,"TR",{});var V=k(n);i=y(V,"TH",{"data-svelte-h":!0}),I(i)!=="svelte-3jponc"&&(i.textContent=T),w=z(V),p=y(V,"TH",{class:!0});var B=k(p);C=K(B,f),B.forEach(h),V.forEach(h),H.forEach(h),g.forEach(h),S.forEach(h),this.h()},h(){j(c,"class","cursor-pointer"),j(p,"class","text-right"),j(a,"class","table table-hover"),j(e,"class","table-container")},m(s,S){O(s,e,S),d(e,a),d(a,t),d(a,m),d(a,c);for(let g=0;g<_.length;g+=1)_[g]&&_[g].m(c,null);d(a,v),d(a,o),d(o,n),d(n,i),d(n,w),d(n,p),d(p,C)},p(s,S){var g;if(S&1){E=F(s[0].data);let u;for(u=0;u<E.length;u+=1){const H=G(s,E,u);_[u]?_[u].p(H,S):(_[u]=P(H),_[u].c(),_[u].m(c,null))}for(;u<_.length;u+=1)_[u].d(1);_.length=E.length}S&1&&f!==(f=((g=s[0].data)==null?void 0:g.length)+"")&&L(C,f)},d(s){s&&h(e),W(_,s)}}}function P(l){let e,a,t=l[2].name+"",r,m,c,v=l[0].params.vote+"",o,n,i,T;function w(){return l[1](l[2])}return{c(){e=b("tr"),a=b("td"),r=x(t),m=D(),c=b("td"),o=x(v),n=D(),this.h()},l(p){e=y(p,"TR",{});var f=k(e);a=y(f,"TD",{});var C=k(a);r=K(C,t),C.forEach(h),m=z(f),c=y(f,"TD",{class:!0});var E=k(c);o=K(E,v),E.forEach(h),n=z(f),f.forEach(h),this.h()},h(){j(c,"class","text-right")},m(p,f){O(p,e,f),d(e,a),d(a,r),d(e,m),d(e,c),d(c,o),d(e,n),i||(T=X(e,"click",w),i=!0)},p(p,f){l=p,f&1&&t!==(t=l[2].name+"")&&L(r,t),f&1&&v!==(v=l[0].params.vote+"")&&L(o,v)},d(p){p&&h(e),i=!1,T()}}}function ae(l){let e,a,t,r,m;document.title=e="Spieler mit Schätzung "+l[0].params.vote+" in Game "+l[0].params.id+" | Killspiel";function c(n,i){if(i&1&&(m=null),m==null&&(m=!!Array.isArray(n[0].data)),m)return te;if(n[0].data===null)return ee}let v=c(l,-1),o=v&&v(l);return{c(){a=D(),t=b("div"),r=b("div"),o&&o.c(),this.h()},l(n){N("svelte-1mnrhsh",document.head).forEach(h),a=z(n),t=y(n,"DIV",{class:!0});var T=k(t);r=y(T,"DIV",{class:!0});var w=k(r);o&&o.l(w),w.forEach(h),T.forEach(h),this.h()},h(){j(r,"class","card mt-6 m-2 w-2/3 mx-auto"),j(t,"class","container mx-auto w-2/5")},m(n,i){O(n,a,i),O(n,t,i),d(t,r),o&&o.m(r,null)},p(n,[i]){i&1&&e!==(e="Spieler mit Schätzung "+n[0].params.vote+" in Game "+n[0].params.id+" | Killspiel")&&(document.title=e),v===(v=c(n,i))&&o?o.p(n,i):(o&&o.d(1),o=v&&v(n),o&&(o.c(),o.m(r,null)))},i:A,o:A,d(n){n&&(h(a),h(t)),o&&o.d()}}}function ne(l,e,a){let{data:t}=e;const r=m=>Q("/user/"+m.id.toString());return l.$$set=m=>{"data"in m&&a(0,t=m.data)},[t,r]}class de extends Y{constructor(e){super(),J(this,e,ne,ae,R,{data:0})}}export{de as component,se as universal};