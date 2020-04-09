import * as React from "react";
import { ISiteMapNode } from "rfluxx-routing";

import { ContainerFactory } from "./ContainerFactory";
import { HomePage } from "./HomePage";


/**
 * The site map node for this page.
 */
export const siteMapNode: ISiteMapNode = {
    caption: "Home Page",
    routeExpression: "/",
    containerFactory: new ContainerFactory(),
    render: p => <HomePage />
};
