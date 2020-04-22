import { IContainer, IContainerBuilder, registerStore, resolveStore } from "rfluxx";
import { IGlobalComponents, ISiteMapNodeContainerBuilder, RouteParameters, SiteMapNodeContainerFactoryBase } from "rfluxx-routing";

import { JobLogSourceStore } from "./JobLogSourceStore";

export class ContainerFactory extends SiteMapNodeContainerFactoryBase
{
    protected registerStores(builder: ISiteMapNodeContainerBuilder): void
    {
        registerStore(builder, "IJobLogSourceStore", (c, injOpt) => JobLogSourceStore(injOpt({
            routeParameters: c.resolve("RouteParametersStream")
        })));
    }
}