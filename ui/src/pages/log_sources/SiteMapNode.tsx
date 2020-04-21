import * as React from "react";
import { ISiteMapNode } from "rfluxx-routing";

import { ContainerFactory } from "./ContainerFactory";
import { LogSourcesPage } from "./LogSourcesPage";


/**
 * The site map node for this page.
 */
export const siteMapNode: ISiteMapNode = {
    caption: "Log Sources",
    routeExpression: "/log_sources",
    containerFactory: new ContainerFactory(),
    render: p => <LogSourcesPage />
};
