"use strict";(self.webpackChunktracetest_docs=self.webpackChunktracetest_docs||[]).push([[7239],{3905:(t,e,r)=>{r.d(e,{Zo:()=>u,kt:()=>f});var n=r(7294);function a(t,e,r){return e in t?Object.defineProperty(t,e,{value:r,enumerable:!0,configurable:!0,writable:!0}):t[e]=r,t}function o(t,e){var r=Object.keys(t);if(Object.getOwnPropertySymbols){var n=Object.getOwnPropertySymbols(t);e&&(n=n.filter((function(e){return Object.getOwnPropertyDescriptor(t,e).enumerable}))),r.push.apply(r,n)}return r}function c(t){for(var e=1;e<arguments.length;e++){var r=null!=arguments[e]?arguments[e]:{};e%2?o(Object(r),!0).forEach((function(e){a(t,e,r[e])})):Object.getOwnPropertyDescriptors?Object.defineProperties(t,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(e){Object.defineProperty(t,e,Object.getOwnPropertyDescriptor(r,e))}))}return t}function i(t,e){if(null==t)return{};var r,n,a=function(t,e){if(null==t)return{};var r,n,a={},o=Object.keys(t);for(n=0;n<o.length;n++)r=o[n],e.indexOf(r)>=0||(a[r]=t[r]);return a}(t,e);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(t);for(n=0;n<o.length;n++)r=o[n],e.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(t,r)&&(a[r]=t[r])}return a}var s=n.createContext({}),l=function(t){var e=n.useContext(s),r=e;return t&&(r="function"==typeof t?t(e):c(c({},e),t)),r},u=function(t){var e=l(t.components);return n.createElement(s.Provider,{value:e},t.children)},p={inlineCode:"code",wrapper:function(t){var e=t.children;return n.createElement(n.Fragment,{},e)}},d=n.forwardRef((function(t,e){var r=t.components,a=t.mdxType,o=t.originalType,s=t.parentName,u=i(t,["components","mdxType","originalType","parentName"]),d=l(r),f=a,k=d["".concat(s,".").concat(f)]||d[f]||p[f]||o;return r?n.createElement(k,c(c({ref:e},u),{},{components:r})):n.createElement(k,c({ref:e},u))}));function f(t,e){var r=arguments,a=e&&e.mdxType;if("string"==typeof t||a){var o=r.length,c=new Array(o);c[0]=d;var i={};for(var s in e)hasOwnProperty.call(e,s)&&(i[s]=e[s]);i.originalType=t,i.mdxType="string"==typeof t?t:a,c[1]=i;for(var l=2;l<o;l++)c[l]=r[l];return n.createElement.apply(null,c)}return n.createElement.apply(null,r)}d.displayName="MDXCreateElement"},4181:(t,e,r)=>{r.r(e),r.d(e,{assets:()=>s,contentTitle:()=>c,default:()=>p,frontMatter:()=>o,metadata:()=>i,toc:()=>l});var n=r(7462),a=(r(7294),r(3905));const o={},c="Quick Start",i={unversionedId:"quick-start",id:"quick-start",title:"Quick Start",description:"In this section, you will:",source:"@site/docs/quick-start.md",sourceDirName:".",slug:"/quick-start",permalink:"/quick-start",draft:!1,editUrl:"https://github.com/kubeshop/tracetest/blob/main/docs/docs/quick-start.md",tags:[],version:"current",frontMatter:{}},s={},l=[{value:"<strong>1. Install the Tracetest CLI</strong>",id:"1-install-the-tracetest-cli",level:3},{value:"<strong>2. Scaffold Tracetest Docker config with the CLI</strong>",id:"2-scaffold-tracetest-docker-config-with-the-cli",level:3},{value:"<strong>3. Point Tracetest to a trace back end</strong>",id:"3-point-tracetest-to-a-trace-back-end",level:3},{value:"<strong>4. Run Tracetest with Docker or the CLI</strong>",id:"4-run-tracetest-with-docker-or-the-cli",level:3},{value:"<strong>5. Create and run a test</strong>",id:"5-create-and-run-a-test",level:3},{value:"<strong>6. View trace and set assertions</strong>",id:"6-view-trace-and-set-assertions",level:3},{value:"Next Steps",id:"next-steps",level:2}],u={toc:l};function p(t){let{components:e,...r}=t;return(0,a.kt)("wrapper",(0,n.Z)({},u,r,{components:e,mdxType:"MDXLayout"}),(0,a.kt)("h1",{id:"quick-start"},"Quick Start"),(0,a.kt)("p",null,"In this section, you will:"),(0,a.kt)("ol",null,(0,a.kt)("li",{parentName:"ol"},"..."),(0,a.kt)("li",{parentName:"ol"},"..."),(0,a.kt)("li",{parentName:"ol"},"..."),(0,a.kt)("li",{parentName:"ol"},"...")),(0,a.kt)("h3",{id:"1-install-the-tracetest-cli"},(0,a.kt)("strong",{parentName:"h3"},"1. Install the Tracetest CLI")),(0,a.kt)("p",null,"..."),(0,a.kt)("h3",{id:"2-scaffold-tracetest-docker-config-with-the-cli"},(0,a.kt)("strong",{parentName:"h3"},"2. Scaffold Tracetest Docker config with the CLI")),(0,a.kt)("p",null,"..."),(0,a.kt)("h3",{id:"3-point-tracetest-to-a-trace-back-end"},(0,a.kt)("strong",{parentName:"h3"},"3. Point Tracetest to a trace back end")),(0,a.kt)("p",null,"..."),(0,a.kt)("h3",{id:"4-run-tracetest-with-docker-or-the-cli"},(0,a.kt)("strong",{parentName:"h3"},"4. Run Tracetest with Docker or the CLI")),(0,a.kt)("p",null,"..."),(0,a.kt)("h3",{id:"5-create-and-run-a-test"},(0,a.kt)("strong",{parentName:"h3"},"5. Create and run a test")),(0,a.kt)("p",null,"..."),(0,a.kt)("h3",{id:"6-view-trace-and-set-assertions"},(0,a.kt)("strong",{parentName:"h3"},"6. View trace and set assertions")),(0,a.kt)("p",null,"..."),(0,a.kt)("h2",{id:"next-steps"},"Next Steps"),(0,a.kt)("p",null,"..."),(0,a.kt)("p",null,"And, if you want, connect with us on ",(0,a.kt)("a",{parentName:"p",href:"https://discord.gg/6zupCZFQbe"},"Discord")," to tell us about your experience!"))}p.isMDXComponent=!0}}]);