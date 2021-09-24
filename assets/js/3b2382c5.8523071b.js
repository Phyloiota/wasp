(self.webpackChunkdoc_ops=self.webpackChunkdoc_ops||[]).push([[1170],{2964:function(e,t,n){"use strict";n.r(t),n.d(t,{frontMatter:function(){return s},contentTitle:function(){return c},metadata:function(){return l},toc:function(){return p},default:function(){return h}});var a=n(2122),r=n(9756),o=(n(7294),n(3905)),i=["components"],s={},c="Structure of the smart contract",l={unversionedId:"tutorial/05",id:"tutorial/05",isDocsHomePage:!1,title:"Structure of the smart contract",description:"Smart contracts are programs, immutably stored in the chain. In the previous",source:"@site/docs/tutorial/05.md",sourceDirName:"tutorial",slug:"/tutorial/05",permalink:"/docs/tutorial/05",editUrl:"https://github.com/iotaledger/wasp/tree/master/documentation/docs/tutorial/05.md",version:"current",frontMatter:{},sidebar:"tutorialSidebar",previous:{title:"Deploying and running a Rust smart contract",permalink:"/docs/tutorial/04"},next:{title:"Invoking smart contracts. Sending a request `on-ledger`",permalink:"/docs/tutorial/06"}},p=[{value:"State",id:"state",children:[]},{value:"Entry points",id:"entry-points",children:[]},{value:"Panic. Exception handling",id:"panic-exception-handling",children:[]}],u={toc:p};function h(e){var t=e.components,s=(0,r.Z)(e,i);return(0,o.kt)("wrapper",(0,a.Z)({},u,s,{components:t,mdxType:"MDXLayout"}),(0,o.kt)("h1",{id:"structure-of-the-smart-contract"},"Structure of the smart contract"),(0,o.kt)("p",null,"Smart contracts are programs, immutably stored in the chain. In the previous\nexample the binary file with the code of the smart contract\n",(0,o.kt)("em",{parentName:"p"},"example_tutorial_bg.wasm")," will be immutably stored in the chain state."),(0,o.kt)("p",null,"The logical structure of an ISCP smart contract is independent of the VM type we\nuse, be it a ",(0,o.kt)("em",{parentName:"p"},"Wasm")," smart contract or any other VM type."),(0,o.kt)("p",null,(0,o.kt)("img",{src:n(4570).Z})),(0,o.kt)("p",null,"Each smart contract on the chain is identified by its name hashed into 4 bytes\nand interpreted as ",(0,o.kt)("inlineCode",{parentName:"p"},"uint32")," value: the so called ",(0,o.kt)("inlineCode",{parentName:"p"},"hname"),". For example,\nthe ",(0,o.kt)("inlineCode",{parentName:"p"},"hname")," of the root contract is ",(0,o.kt)("inlineCode",{parentName:"p"},"0xcebf5908"),", the unique identifier of the\n",(0,o.kt)("inlineCode",{parentName:"p"},"root")," contract in every chain. The exception is ",(0,o.kt)("inlineCode",{parentName:"p"},"_default")," contract which always has hname ",(0,o.kt)("inlineCode",{parentName:"p"},"0x00000000"),"."),(0,o.kt)("p",null,"Each smart contract instance has a program with a collection of entry points and\na state. An entry point is a function of the program through which the program\ncan be invoked. The ",(0,o.kt)("inlineCode",{parentName:"p"},"example1")," contract above has three entry\npoints: ",(0,o.kt)("inlineCode",{parentName:"p"},"storeString"),", ",(0,o.kt)("inlineCode",{parentName:"p"},"getString")," and ",(0,o.kt)("inlineCode",{parentName:"p"},"withdrawIota"),"."),(0,o.kt)("p",null,"There are several ways to invoke an entry point: a request, a call and a view\ncall, depending on the type of the entry point."),(0,o.kt)("p",null,"The smart contract program can access its state and account through an interface\nlayer called the ",(0,o.kt)("em",{parentName:"p"},"Sandbox"),"."),(0,o.kt)("h2",{id:"state"},"State"),(0,o.kt)("p",null,"The smart contract state is its data, with each update stored on the chain. The\nstate can only be modified by the smart contract program itself. There are two\nparts of the state:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},"A collection of key/value pairs called the ",(0,o.kt)("inlineCode",{parentName:"li"},"data state"),". Each key and value\nare byte arrays of arbitrary size (there are practical limits set by the\ndatabase, of course). The value of the key/value pair is always retrieved by\nits key."),(0,o.kt)("li",{parentName:"ul"},"A collection of ",(0,o.kt)("inlineCode",{parentName:"li"},"color: balance")," pairs called the ",(0,o.kt)("inlineCode",{parentName:"li"},"account"),". The account\nrepresents the balances of tokens of specific colors controlled by the smart\ncontract. Receiving and spending tokens into/from the account means changing\nthe account's balances.")),(0,o.kt)("p",null,"Only the smart contract program can change its data state and spend from its\naccount. Tokens can be sent to the smart contract account by any other agent on\nthe ledger, be it a wallet with an address or another smart contract."),(0,o.kt)("p",null,"See ",(0,o.kt)("a",{parentName:"p",href:"/docs/contract_core/accounts"},"Accounts")," for more info on sending and receiving tokens."),(0,o.kt)("h2",{id:"entry-points"},"Entry points"),(0,o.kt)("p",null,"There are two types of entry points:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("em",{parentName:"li"},"Full entry points")," or just ",(0,o.kt)("em",{parentName:"li"},"entry points"),". These functions can modify\n(mutate) the state of the smart contract."),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("em",{parentName:"li"},"View entry points")," or ",(0,o.kt)("em",{parentName:"li"},"views"),". These are read-only functions. They are used\nonly to retrieve the information from the smart contract state. They can\u2019t\nmodify the state, i.e. are read-only calls.")),(0,o.kt)("p",null,"The ",(0,o.kt)("inlineCode",{parentName:"p"},"example1")," program has three entry points:"),(0,o.kt)("ul",null,(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("inlineCode",{parentName:"p"},"storeString")," a full entry point. It first checks if parameter\ncalled ",(0,o.kt)("inlineCode",{parentName:"p"},"paramString")," exist. If so, it stores the string value of the parameter\ninto the state variable ",(0,o.kt)("inlineCode",{parentName:"p"},"storedString"),". If parameter ",(0,o.kt)("inlineCode",{parentName:"p"},"paramString")," is missing,\nthe program panics.")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("inlineCode",{parentName:"p"},"getString")," is a view entry point that returns the value of the\nvariable ",(0,o.kt)("inlineCode",{parentName:"p"},"storedString"),".")),(0,o.kt)("li",{parentName:"ul"},(0,o.kt)("p",{parentName:"li"},(0,o.kt)("inlineCode",{parentName:"p"},"withdrawIota")," is a full entry point that checks if the caller is and address\nand if the caller is equal to the creator of smart contract. If not, it\npanics. If it passes the validation, the program sends all the iotas contained\nin the smart contract's account to the caller."))),(0,o.kt)("p",null,"Note that in the ",(0,o.kt)("inlineCode",{parentName:"p"},"example1")," the Rust functions associated with full entry points\ntake a parameter of type ",(0,o.kt)("inlineCode",{parentName:"p"},"ScFuncContext"),". It gives full (read-write) access to\nthe state. In contrast, ",(0,o.kt)("inlineCode",{parentName:"p"},"getString")," is a view entry point and its associated\nfunction parameter has type ",(0,o.kt)("inlineCode",{parentName:"p"},"ScViewContext"),". A view is not allowed to mutate\nthe state."),(0,o.kt)("h2",{id:"panic-exception-handling"},"Panic. Exception handling"),(0,o.kt)("p",null,"The following test posts a request to the ",(0,o.kt)("inlineCode",{parentName:"p"},"example1")," smart contract without\nthe expected parameter ",(0,o.kt)("inlineCode",{parentName:"p"},"paramString"),". The\nstatement ",(0,o.kt)("inlineCode",{parentName:"p"},'ctx.require(par.exists(), "string parameter not found");')," makes\nthe smart contract panic if the condition is not satisfied."),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre",className:"language-go"},'func TestTutorial4(t *testing.T) {\n    env := solo.New(t, false, false, seed)\n    chain := env.NewChain(nil, "ex4")\n    // deploy the contract on chain\n    err := chain.DeployWasmContract(nil, "example1", "example_tutorial_bg.wasm")\n    require.NoError(t, err)\n\n    // call contract incorrectly (omit \'paramString\')\n    req := solo.NewCallParams("example1", "storeString").WithIotas(1)\n    _, err = chain.PostRequestSync(req, nil)\n    require.Error(t, err)\n    t.Logf("error: \'%v\'", err)\n}\n')),(0,o.kt)("p",null,"The fragments in the output of the test:"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"37:34.189474700 PANIC   TestTutorial4.ex4   vmcontext/log.go:12 string parameter not found\n\n37:34.192828900 INFO    TestTutorial4.ex4   solo/run.go:148 REQ: 'tx/[0]9r5zoeusdwTcWkDTEMYjeqNj8reiUsLiHF81vExPrvNW: 'panic in VM: string parameter not found''\n")),(0,o.kt)("p",null,"It shows that the panic indeed occurred. The test passes because the resulting\nerror was expected."),(0,o.kt)("p",null,"The log record"),(0,o.kt)("pre",null,(0,o.kt)("code",{parentName:"pre"},"37:34.192828900 INFO    TestTutorial4.ex4   solo/run.go:148 REQ: 'tx/[0]9r5zoeusdwTcWkDTEMYjeqNj8reiUsLiHF81vExPrvNW: 'panic in VM: string parameter not found''\n")),(0,o.kt)("p",null,"is a printed receipt of the request. It is stored on the chain for each request processed."),(0,o.kt)("p",null,"Note that this test ends with the state ",(0,o.kt)("inlineCode",{parentName:"p"},"#4"),", despite the fact that the last\nrequest to the smart contract failed. This is important: ",(0,o.kt)("strong",{parentName:"p"},"whatever happens\nduring the execution of a smart contract's full entry point, processing of the\nrequest always results in the state transition"),"."),(0,o.kt)("p",null,"The VM context catches exceptions (panics) in the program. Its consequences are\nrecorded in the state of the chain during the fallback processing, no matter if\nthe panic was triggered by the logic of the smart contract or whether it was\ntriggered by the sandbox run-time code."),(0,o.kt)("p",null,"In the case of ",(0,o.kt)("inlineCode",{parentName:"p"},"example1")," the error event was recorded in the immutable record\nlog of the chain, aka ",(0,o.kt)("inlineCode",{parentName:"p"},"receipt"),", but the data state of the smart contract wasn't modified. In\nother cases the fallback actions may be more complex."))}h.isMDXComponent=!0},3905:function(e,t,n){"use strict";n.d(t,{Zo:function(){return p},kt:function(){return m}});var a=n(7294);function r(e,t,n){return t in e?Object.defineProperty(e,t,{value:n,enumerable:!0,configurable:!0,writable:!0}):e[t]=n,e}function o(e,t){var n=Object.keys(e);if(Object.getOwnPropertySymbols){var a=Object.getOwnPropertySymbols(e);t&&(a=a.filter((function(t){return Object.getOwnPropertyDescriptor(e,t).enumerable}))),n.push.apply(n,a)}return n}function i(e){for(var t=1;t<arguments.length;t++){var n=null!=arguments[t]?arguments[t]:{};t%2?o(Object(n),!0).forEach((function(t){r(e,t,n[t])})):Object.getOwnPropertyDescriptors?Object.defineProperties(e,Object.getOwnPropertyDescriptors(n)):o(Object(n)).forEach((function(t){Object.defineProperty(e,t,Object.getOwnPropertyDescriptor(n,t))}))}return e}function s(e,t){if(null==e)return{};var n,a,r=function(e,t){if(null==e)return{};var n,a,r={},o=Object.keys(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||(r[n]=e[n]);return r}(e,t);if(Object.getOwnPropertySymbols){var o=Object.getOwnPropertySymbols(e);for(a=0;a<o.length;a++)n=o[a],t.indexOf(n)>=0||Object.prototype.propertyIsEnumerable.call(e,n)&&(r[n]=e[n])}return r}var c=a.createContext({}),l=function(e){var t=a.useContext(c),n=t;return e&&(n="function"==typeof e?e(t):i(i({},t),e)),n},p=function(e){var t=l(e.components);return a.createElement(c.Provider,{value:t},e.children)},u={inlineCode:"code",wrapper:function(e){var t=e.children;return a.createElement(a.Fragment,{},t)}},h=a.forwardRef((function(e,t){var n=e.components,r=e.mdxType,o=e.originalType,c=e.parentName,p=s(e,["components","mdxType","originalType","parentName"]),h=l(n),m=r,d=h["".concat(c,".").concat(m)]||h[m]||u[m]||o;return n?a.createElement(d,i(i({ref:t},p),{},{components:n})):a.createElement(d,i({ref:t},p))}));function m(e,t){var n=arguments,r=t&&t.mdxType;if("string"==typeof e||r){var o=n.length,i=new Array(o);i[0]=h;var s={};for(var c in t)hasOwnProperty.call(t,c)&&(s[c]=t[c]);s.originalType=e,s.mdxType="string"==typeof e?e:r,i[1]=s;for(var l=2;l<o;l++)i[l]=n[l];return a.createElement.apply(null,i)}return a.createElement.apply(null,n)}h.displayName="MDXCreateElement"},4570:function(e,t,n){"use strict";t.Z=n.p+"assets/images/SC-structure-95a09ea794362efb20ff6a4d06627285.png"}}]);