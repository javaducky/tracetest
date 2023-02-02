"use strict";(self.webpackChunktracetest_docs=self.webpackChunktracetest_docs||[]).push([[1440],{3905:(e,t,n)=>{n.d(t,{Zo:()=>p,kt:()=>f});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function s(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?s(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):s(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function a(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},s=Object.keys(e);for(r=0;r<s.length;r++)n=s[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var s=Object.getOwnPropertySymbols(e);for(r=0;r<s.length;r++)n=s[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var l=r.createContext({}),c=function(e){var t=r.useContext(l),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},p=function(e){var t=c(e.components);return r.createElement(l.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},d=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,s=e.originalType,l=e.parentName,p=a(e,["components","mdxType","originalType","parentName"]),d=c(n),f=o,h=d["".concat(l,".").concat(f)]||d[f]||u[f]||s;return n?r.createElement(h,i(i({ref:t},p),{},{components:n})):r.createElement(h,i({ref:t},p))}));function f(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var s=n.length,i=new Array(s);i[0]=d;var a={};for(var l in t)hasOwnProperty.call(t,l)&&(a[l]=t[l]);a.originalType=e,a.mdxType="string"==typeof e?e:o,i[1]=a;for(var c=2;c<s;c++)i[c]=n[c];return r.createElement.apply(null,i)}return r.createElement.apply(null,n)}d.displayName="MDXCreateElement"},7403:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>i,default:()=>u,frontMatter:()=>s,metadata:()=>a,toc:()=>c});var r=n(7462),o=(n(7294),n(3905));const s={},i="Edit a Test",a={unversionedId:"run-exports",id:"run-exports",title:"Edit a Test",description:"Tracetest allows you to export the different set of information displayed for assertions and checks in a way you can use it as input for other tools, create text based tests to use on your CI/CD pipelines using the CLI and more options.",source:"@site/docs/run-exports.md",sourceDirName:".",slug:"/run-exports",permalink:"/run-exports",draft:!1,editUrl:"https://github.com/kubeshop/tracetest/blob/main/docs/docs/run-exports.md",tags:[],version:"current",frontMatter:{}},l={},c=[{value:"JUnit Results XML",id:"junit-results-xml",level:2},{value:"Test Definition YAML",id:"test-definition-yaml",level:2}],p={toc:c};function u(e){let{components:t,...s}=e;return(0,o.kt)("wrapper",(0,r.Z)({},p,s,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"edit-a-test"},"Edit a Test"),(0,o.kt)("p",null,"Tracetest allows you to export the different set of information displayed for assertions and checks in a way you can use it as input for other tools, create text based tests to use on your CI/CD pipelines using the CLI and more options."),(0,o.kt)("p",null,"The current supported exports are:"),(0,o.kt)("ol",null,(0,o.kt)("li",{parentName:"ol"},"JUnit results XML."),(0,o.kt)("li",{parentName:"ol"},"Test Definition YAML.")),(0,o.kt)("p",null,"To access any of the available exports, go to the run/trace page details for any test and, at the header level, you'll find a three dot menu which will display the options."),(0,o.kt)("p",null,(0,o.kt)("img",{alt:"Export Trace Options",src:n(9639).Z,width:"3354",height:"850"})),(0,o.kt)("h2",{id:"junit-results-xml"},"JUnit Results XML"),(0,o.kt)("p",null,"To access the JUnit XML file, select the JUnit option from the dropdown and you'll find the file viewer modal with the location to download the file.\nThe JUnit report contains the results from each of the assertions added to the test and their statuses. Depending on how many assertions the test has, this file will grow."),(0,o.kt)("p",null,(0,o.kt)("img",{alt:"Export Trace JUnit",src:n(84).Z,width:"2474",height:"1248"})),(0,o.kt)("h2",{id:"test-definition-yaml"},"Test Definition YAML"),(0,o.kt)("p",null,"The Tracetest CLI allows you to execute text based tests. This means you can store all of your tests in a repo, keep track of the different versions and use them for your CI/CD process.\nAn easy way to start is to export the test definition directly from the UI by selecting the option from the dropdown.\nThe file viewer modal will popup where you can copy paste or download the file."),(0,o.kt)("p",null,(0,o.kt)("img",{alt:"Export Trace Test Definition",src:n(3070).Z,width:"2518",height:"1422"})))}u.isMDXComponent=!0},84:(e,t,n)=>{n.d(t,{Z:()=>r});const r=n.p+"assets/images/exports-junit-b9e427516121a5310af5cbec3d8bf780.png"},3070:(e,t,n)=>{n.d(t,{Z:()=>r});const r=n.p+"assets/images/exports-test-definition-17e5df5e4f432d6096272a6406944556.png"},9639:(e,t,n)=>{n.d(t,{Z:()=>r});const r=n.p+"assets/images/exports-trace-options-94e1015a6002079400ab91cd6903a578.png"}}]);