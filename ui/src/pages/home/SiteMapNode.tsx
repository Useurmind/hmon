import * as React from "react";
import { ISiteMapNode } from "rfluxx-routing";

import { ContainerFactory } from "./ContainerFactory";
import { HomePage } from "./HomePage";

import { siteMapNode as logSourcePageNode } from "../log_sources/SiteMapNode";
import { siteMapNode as jobLogSourcePageNode } from "../job_log_source/SiteMapNode";


/**
 * The site map node for this page.
 */
export const siteMapNode: ISiteMapNode = {
    caption: "Home Page",
    routeExpression: "/",
    containerFactory: new ContainerFactory(),
    render: p => <HomePage />,
    children: [
        logSourcePageNode,
        jobLogSourcePageNode
    ]
};
