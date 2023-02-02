"use strict";(self.webpackChunktracetest_docs=self.webpackChunktracetest_docs||[]).push([[8517],{3905:(e,t,r)=>{r.d(t,{Zo:()=>p,kt:()=>d});var a=r(7294);function n(e,t,r){return t in e?Object.defineProperty(e,t,{value:r,enumerable:!0,configurable:!0,writable:!0}):e[t]=r,e}function o(e,t){var r=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),r.push.apply(r,a)}return r}function i(e){for(var t=1;t<arguments.length;t++){var r=null!=arguments[t]?arguments[t]:{};t%2?o(Object(r),!0).forEach((function(t){n(e,t,r[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(r)):o(Object(r)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(r,t))}))}return e}function s(e,t){if(null==e)return{};var r,a,n=function(e,t){if(null==e)return{};var r,a,n={},o=Object.keys(e);for(a=0;a<o.length;a++)r=o[a],t.indexOf(r)>=0||(n[r]=e[r]);return n}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)r=o[a],t.indexOf(r)>=0||Object.prototype.propertyIsEnumerable.call(e,r)&&(n[r]=e[r])}return n}var l=a.createContext({}),c=function(e){var t=a.useContext(l),r=t;return e&&(r="function"==typeof e?e(t):i(i({},t),e)),r},p=function(e){var t=c(e.components);return a.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},m=a.forwardRef((function(e,t){var r=e.components,n=e.mdxType,o=e.originalType,l=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),m=c(r),d=n,h=m["".concat(l,".").concat(d)]||m[d]||u[d]||o;return r?a.createElement(h,i(i({ref:t},p),{},{components:r})):a.createElement(h,i({ref:t},p))}));function d(e,t){var r=arguments,n=t&&t.mdxType;if("string"==typeof e||n){var o=r.length,i=new Array(o);i[0]=m;var s={};for(var l in t)hasOwnProperty.call(t,l)&&(s[l]=t[l]);s.originalType=e,s.mdxType="string"==typeof e?e:n,i[1]=s;for(var c=2;c<o;c++)i[c]=r[c];return a.createElement.apply(null,i)}return a.createElement.apply(null,r)}m.displayName="MDXCreateElement"},2427:(e,t,r)=>{r.r(t),r.d(t,{assets:()=>l,contentTitle:()=>i,default:()=>u,frontMatter:()=>o,metadata:()=>s,toc:()=>c});var a=r(7462),n=(r(7294),r(3905));const o={},i="OTel-Demo testing game @ KubeconNA 2022",s={unversionedId:"kubecon",id:"kubecon",title:"OTel-Demo testing game @ KubeconNA 2022",description:"Play our game, learn to run trace-based tests using TraceTest, get familiar with the OpenTelemetry Community Demo, and win a Plushie!",source:"@site/docs/kubecon.md",sourceDirName:".",slug:"/kubecon",permalink:"/kubecon",draft:!1,editUrl:"https://github.com/kubeshop/tracetest/blob/main/docs/docs/kubecon.md",tags:[],version:"current",frontMatter:{}},l={},c=[{value:"Instructions",id:"instructions",level:2},{value:"Help",id:"help",level:2}],p={toc:c};function u(e){let{components:t,...r}=e;return(0,n.kt)("wrapper",(0,a.Z)({},p,r,{components:t,mdxType:"MDXLayout"}),(0,n.kt)("h1",{id:"otel-demo-testing-game--kubeconna-2022"},"OTel-Demo testing game @ KubeconNA 2022"),(0,n.kt)("p",null,"Play our game, learn to run trace-based tests using TraceTest, get familiar with the ",(0,n.kt)("a",{parentName:"p",href:"https://github.com/open-telemetry/opentelemetry-demo"},"OpenTelemetry Community Demo"),", and win a Plushie!"),(0,n.kt)("p",null,"We have deployed the otel demo at ",(0,n.kt)("a",{parentName:"p",href:"http://otel-demo.tracetest.io/"},"http://otel-demo.tracetest.io/"),", and a Tracetest instance ready to connect to it at ",(0,n.kt)("a",{parentName:"p",href:"http://tracetest-otel-demo.tracetest.io/"},"http://tracetest-otel-demo.tracetest.io/")),(0,n.kt)("h2",{id:"instructions"},"Instructions"),(0,n.kt)("ol",null,(0,n.kt)("li",{parentName:"ol"},(0,n.kt)("p",{parentName:"li"},"Go to ",(0,n.kt)("a",{parentName:"p",href:"http://tracetest-otel-demo.tracetest.io/"},"http://tracetest-otel-demo.tracetest.io/"))),(0,n.kt)("li",{parentName:"ol"},(0,n.kt)("p",{parentName:"li"},"Create a new test. The goal will be to run tests over a trace generated by a call to the ",(0,n.kt)("inlineCode",{parentName:"p"},"Checkout API"),". You can use the ",(0,n.kt)("inlineCode",{parentName:"p"},"Choose example")," dropdown to select the correct one. Add your name to the name of the test.")),(0,n.kt)("li",{parentName:"ol"},(0,n.kt)("p",{parentName:"li"},"When the test finishes running, you will be able to see and play around with the generated trace. Now comes the interesting part.")),(0,n.kt)("li",{parentName:"ol"},(0,n.kt)("p",{parentName:"li"},"Using the ",(0,n.kt)("inlineCode",{parentName:"p"},"Test")," section, find your way into adding ",(0,n.kt)("inlineCode",{parentName:"p"},"Test Specs")," to do the following assertions:"),(0,n.kt)("ul",{parentName:"li"},(0,n.kt)("li",{parentName:"ul"},"All spans with a type of ",(0,n.kt)("inlineCode",{parentName:"li"},"RPC")," have been successful (a.k.a ",(0,n.kt)("inlineCode",{parentName:"li"},"grpc.status_code")," equals 0)"),(0,n.kt)("li",{parentName:"ul"},"Since we\u2019re trying to checkout an empty cart, the shipping cost will be zero. Assert that the ",(0,n.kt)("inlineCode",{parentName:"li"},"get-quote")," span reflects that information"),(0,n.kt)("li",{parentName:"ul"},"Finally, we want to be sure the user provided credit card is valid, so make sure the ",(0,n.kt)("inlineCode",{parentName:"li"},"charge")," span shows if the given card is valid or not.")))),(0,n.kt)("ol",{start:5},(0,n.kt)("li",{parentName:"ol"},(0,n.kt)("p",{parentName:"li"},'Don\u2019t forget to "Publish" your changes once you created all your specs.')),(0,n.kt)("li",{parentName:"ol"},(0,n.kt)("p",{parentName:"li"},"Complete ",(0,n.kt)("a",{parentName:"p",href:"https://forms.gle/pAyCFjKUeBAKhTFP7"},"our survey"),". You will need to copy/paste your test URL. Make sure it is correct by copy/pasting it into a new browser window. If it\u2019s correct, you should see your test with all the specs you added.")),(0,n.kt)("li",{parentName:"ol"},(0,n.kt)("p",{parentName:"li"},"Come by the stand and reclaim your plushie!"))),(0,n.kt)("h2",{id:"help"},"Help"),(0,n.kt)("ul",null,(0,n.kt)("li",{parentName:"ul"},"You can use our docs: ",(0,n.kt)("a",{parentName:"li",href:"https://docs.tracetest.io/"},"https://docs.tracetest.io/"),". Some relevant pages:",(0,n.kt)("ul",{parentName:"li"},(0,n.kt)("li",{parentName:"ul"},"Creating tests: ",(0,n.kt)("a",{parentName:"li",href:"https://docs.tracetest.io/create-test/"},"https://docs.tracetest.io/create-test/")),(0,n.kt)("li",{parentName:"ul"},"Adding test spec: ",(0,n.kt)("a",{parentName:"li",href:"https://docs.tracetest.io/adding-assertions/"},"https://docs.tracetest.io/adding-assertions/")),(0,n.kt)("li",{parentName:"ul"},"Advanced selectors: ",(0,n.kt)("a",{parentName:"li",href:"https://docs.tracetest.io/advanced-selectors/"},"https://docs.tracetest.io/advanced-selectors/")))),(0,n.kt)("li",{parentName:"ul"},"Selectors can be hard, so we have a cheatsheet: ",(0,n.kt)("a",{parentName:"li",href:"https://docs.tracetest.io/img/cheatsheet.pdf"},"https://docs.tracetest.io/img/cheatsheet.pdf")),(0,n.kt)("li",{parentName:"ul"},"You can ask for help in our Discord channel: ",(0,n.kt)("a",{parentName:"li",href:"https://discord.com/invite/6zupCZFQbe"},"https://discord.com/invite/6zupCZFQbe"))))}u.isMDXComponent=!0}}]);