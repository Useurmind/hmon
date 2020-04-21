import { IContainer, IContainerBuilder, registerStore, resolveStore } from "rfluxx";
import { IGlobalComponents, ISiteMapNodeContainerBuilder, RouteParameters, SiteMapNodeContainerFactoryBase } from "rfluxx-routing";

import { LogSourcePageStore } from "./LogSourcePageStore";

export class ContainerFactory extends SiteMapNodeContainerFactoryBase
{
    protected registerStores(builder: ISiteMapNodeContainerBuilder): void
    {
        registerStore(builder, "ILogSourcePageStore", (c, injOpt) => LogSourcePageStore(injOpt({
        })));
    }
}