import{s as A,a as b,f as _,l as f,S as B,d as n,c as q,g as v,h as w,m as x,j as D,i as C,x as s,n as k,y as R,z as F}from"../chunks/scheduler.QGVxIACg.js";import{S as G,i as J}from"../chunks/index.ElInTfSE.js";import{p as L}from"../chunks/stores.WTvSuIuB.js";function M(r){let i,c,e,a,p=r[0].status+"",g,H,m=r[0].error.message+"",E,I,P,K,d,o,V,u=r[0].params.id+"",y,j;return document.title=i=r[0].status+" Player doesn't exist | Killspiel",{c(){c=b(),e=_("div"),a=_("h1"),g=f(p),H=b(),E=f(m),I=b(),P=_("hr"),K=b(),d=_("div"),o=_("p"),V=f("Player with the id "),y=f(u),j=f(" doesn't exist."),this.h()},l(t){B("svelte-1qtgiku",document.head).forEach(n),c=q(t),e=v(t,"DIV",{class:!0});var h=w(e);a=v(h,"H1",{class:!0});var S=w(a);g=x(S,p),H=q(S),E=x(S,m),S.forEach(n),I=q(h),P=v(h,"HR",{class:!0}),K=q(h),d=v(h,"DIV",{class:!0});var z=w(d);o=v(z,"P",{});var $=w(o);V=x($,"Player with the id "),y=x($,u),j=x($," doesn't exist."),$.forEach(n),z.forEach(n),h.forEach(n),this.h()},h(){D(a,"class","text-6xl text-center"),D(P,"class","p-2 mt-3"),D(d,"class","text-center text-2xl"),D(e,"class","container mx-auto p-2")},m(t,l){C(t,c,l),C(t,e,l),s(e,a),s(a,g),s(a,H),s(a,E),s(e,I),s(e,P),s(e,K),s(e,d),s(d,o),s(o,V),s(o,y),s(o,j)},p(t,[l]){l&1&&i!==(i=t[0].status+" Player doesn't exist | Killspiel")&&(document.title=i),l&1&&p!==(p=t[0].status+"")&&k(g,p),l&1&&m!==(m=t[0].error.message+"")&&k(E,m),l&1&&u!==(u=t[0].params.id+"")&&k(y,u)},i:R,o:R,d(t){t&&(n(c),n(e))}}}function N(r,i,c){let e;return F(r,L,a=>c(0,e=a)),[e]}class U extends G{constructor(i){super(),J(this,i,N,M,A,{})}}export{U as component};