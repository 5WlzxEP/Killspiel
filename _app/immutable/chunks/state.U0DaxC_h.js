import{w as r}from"./index.HhSS_fSS.js";import{Y as s}from"./scheduler.QGVxIACg.js";const e=r(0);function u(t){switch(t){case 0:return"Bereit";case 1:return"Sammle Schätzungen";case 2:return"Warte auf Auflösung";case 3:return"Verkünde Ergebnis"}return"Unbekannter Status"}e.subscribe(()=>{s(e)==3&&setTimeout(()=>{e.set(0)},3e3)});export{u as p,e as s};
