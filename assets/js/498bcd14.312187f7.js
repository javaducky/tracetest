"use strict";(self.webpackChunktracetest_docs=self.webpackChunktracetest_docs||[]).push([[9462],{3905:(e,t,n)=>{n.d(t,{Zo:()=>l,kt:()=>d});var r=n(7294);function o(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function a(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var r=Object.getOwnPropertySymbols(e);t&&(r=r.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,r)}return n}function s(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?a(Object(n),!0).forEach((function(t){o(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):a(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function i(e,t){if(null==e)return{};var n,r,o=function(e,t){if(null==e)return{};var n,r,o={},a=Object.keys(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||(o[n]=e[n]);return o}(e,t);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);for(r=0;r<a.length;r++)n=a[r],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(o[n]=e[n])}return o}var c=r.createContext({}),p=function(e){var t=r.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):s(s({},t),e)),n},l=function(e){var t=p(e.components);return r.createElement(c.Provider,{value:t},e.children)},m={inlineCode:"code",wrapper:function(e){var t=e.children;return r.createElement(r.Fragment,{},t)}},u=r.forwardRef((function(e,t){var n=e.components,o=e.mdxType,a=e.originalType,c=e.parentName,l=i(e,["components","mdxType","originalType","parentName"]),u=p(n),d=o,v=u["".concat(c,".").concat(d)]||u[d]||m[d]||a;return n?r.createElement(v,s(s({ref:t},l),{},{components:n})):r.createElement(v,s({ref:t},l))}));function d(e,t){var n=arguments,o=t&&t.mdxType;if("string"==typeof e||o){var a=n.length,s=new Array(a);s[0]=u;var i={};for(var c in t)hasOwnProperty.call(t,c)&&(i[c]=t[c]);i.originalType=e,i.mdxType="string"==typeof e?e:o,s[1]=i;for(var p=2;p<a;p++)s[p]=n[p];return r.createElement.apply(null,s)}return r.createElement.apply(null,n)}u.displayName="MDXCreateElement"},3820:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>c,contentTitle:()=>s,default:()=>m,frontMatter:()=>a,metadata:()=>i,toc:()=>p});var r=n(7462),o=(n(7294),n(3905));const a={},s="Environments",i={unversionedId:"concepts/environments",id:"concepts/environments",title:"Environments",description:"A common use case for tests is to assert the same behavior across multiple environments (dev, staging, and production, for example). To make sure all environments will have the same behavior, it is important that the tests executed against those environments test the same aspects. To reduce the risks of diverging tests, Tracetest allows you to organize different environments configurations using global objects called Environments.",source:"@site/docs/concepts/environments.md",sourceDirName:"concepts",slug:"/concepts/environments",permalink:"/concepts/environments",draft:!1,editUrl:"https://github.com/kubeshop/tracetest/blob/main/docs/docs/concepts/environments.md",tags:[],version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Assertions",permalink:"/concepts/assertions"},next:{title:"Selectors",permalink:"/concepts/selectors"}},c={},p=[{value:"How Environments Work",id:"how-environments-work",level:2}],l={toc:p};function m(e){let{components:t,...n}=e;return(0,o.kt)("wrapper",(0,r.Z)({},l,n,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"environments"},"Environments"),(0,o.kt)("p",null,"A common use case for tests is to assert the same behavior across multiple environments (dev, staging, and production, for example). To make sure all environments will have the same behavior, it is important that the tests executed against those environments test the same aspects. To reduce the risks of diverging tests, Tracetest allows you to organize different environments configurations using global objects called ",(0,o.kt)("strong",{parentName:"p"},"Environments"),"."),(0,o.kt)("h2",{id:"how-environments-work"},"How Environments Work"),(0,o.kt)("p",null,"Environments are objects containing variables that can be referenced by tests. You can use a single test and provide the information on which environment object will be used to run the test. To illustrate this, consider an app that is deployed in three stages: ",(0,o.kt)("inlineCode",{parentName:"p"},"dev"),", ",(0,o.kt)("inlineCode",{parentName:"p"},"staging"),", and ",(0,o.kt)("inlineCode",{parentName:"p"},"production"),". We can execute the same test in all those environments, however, both ",(0,o.kt)("inlineCode",{parentName:"p"},"URL")," and ",(0,o.kt)("inlineCode",{parentName:"p"},"credentials")," change from environment to environment. To run the same test against the three deployments of the app, you can create three environments:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-dotenv"},"# dev.env\nURL=https://app-dev.com\nAPI_TOKEN=dev-key\n")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-dotenv"},"# staging.env\nURL=https://app-staging.com\nAPI_TOKEN=staging-key\n")),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-dotenv"},"# production.env\nURL=https://app-prod.com\nAPI_TOKEN=prod-key\n")),(0,o.kt)("p",null,"Now consider the following test:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-yaml"},'type: Test\nspecs:\n  name: Test user creation\n  trigger:\n    type: http\n    httpRequest:\n        url: "${env:URL}/api/users"\n        method: POST\n        body: \'{}\'\n        authentication:\n          type: bearer\n          bearer:\n            token: "${env:API_TOKEN}"\n')),(0,o.kt)("p",null,"Both ",(0,o.kt)("inlineCode",{parentName:"p"},"env:URL")," and ",(0,o.kt)("inlineCode",{parentName:"p"},"env:API_TOKEN")," would be replaced by the variables defined in the chosen environment where the test will run. So, if the chosen environment was ",(0,o.kt)("inlineCode",{parentName:"p"},"dev.env"),", its values would be replaced by ",(0,o.kt)("inlineCode",{parentName:"p"},"https://app-dev.com")," and ",(0,o.kt)("inlineCode",{parentName:"p"},"dev-key"),", respectively."))}m.isMDXComponent=!0}}]);