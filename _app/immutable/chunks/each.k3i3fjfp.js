import{a as q,t as z}from"./index.ElInTfSE.js";import{L as B}from"./scheduler.QGVxIACg.js";function F(n){return(n==null?void 0:n.length)!==void 0?n:Array.from(n)}function C(n,d){z(n,1,1,()=>{d.delete(n.key)})}function G(n,d){n.f(),C(n,d)}function H(n,d,x,S,A,g,o,L,p,j,a,k){let i=n.length,c=g.length,f=i;const u={};for(;f--;)u[n[f].key]=f;const h=[],w=new Map,m=new Map,_=[];for(f=c;f--;){const e=k(A,g,f),s=x(e);let t=o.get(s);t?S&&_.push(()=>t.p(e,d)):(t=j(s,e),t.c()),w.set(s,h[f]=t),s in u&&m.set(s,Math.abs(f-u[s]))}const M=new Set,v=new Set;function y(e){q(e,1),e.m(L,a),o.set(e.key,e),a=e.first,c--}for(;i&&c;){const e=h[c-1],s=n[i-1],t=e.key,l=s.key;e===s?(a=e.first,i--,c--):w.has(l)?!o.has(t)||M.has(t)?y(e):v.has(l)?i--:m.get(t)>m.get(l)?(v.add(t),y(e)):(M.add(l),i--):(p(s,o),i--)}for(;i--;){const e=n[i];w.has(e.key)||p(e,o)}for(;c;)y(h[c-1]);return B(_),h}export{F as e,G as f,C as o,H as u};
