import { IContainer, IContainerBuilder, registerStore, resolveStore } from "rfluxx";
import { IGlobalComponents, ISiteMapNodeContainerBuilder, RouteParameters, SiteMapNodeContainerFactoryBase } from "rfluxx-routing";

export class ContainerFactory extends SiteMapNodeContainerFactoryBase
{
    protected registerStores(builder: ISiteMapNodeContainerBuilder): void
    {
        // registerStore(builder, "ISelectPageStore", (c, injOpt) => new SelectPageStore(injOpt({
        //     pageStore: c.resolve("IPageStore"),
        //     pageRequest: c.resolve("IPageRequest")
        // })));
    }
}