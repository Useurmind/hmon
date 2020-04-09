import { IContainer, IContainerBuilder, registerStore, resolveStore } from "rfluxx";

import { registerRfluxxDebug } from "rfluxx-debug";


import { registerThemesGlobally } from "rfluxx-mui-theming";

import { GlobalContainerFactoryBase, IGlobalComponents, IGlobalContainerBuilder, RouteParameters } from "rfluxx-routing";

import { registerResourcesGlobally } from "rfluxx-i18n";

import { languages } from "./i18n/Languages";


export class GlobalContainerFactory extends GlobalContainerFactoryBase
{
    protected registerStores(builder: IGlobalContainerBuilder): void
    {
        
        registerThemesGlobally(builder, [ "Default" ]);
        

        
        registerResourcesGlobally(builder, languages);
                

        
        registerRfluxxDebug(builder);
        

        // register your global stores here
        // registerStore(builder, "IFormPageStore", (c, injOpt) => new FormPageStore(injOpt({
        //     pageStore: c.resolve("IPageStore")
        // })));
    }
}
