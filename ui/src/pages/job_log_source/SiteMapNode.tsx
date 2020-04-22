import * as React from "react";
import { ISiteMapNode } from "rfluxx-routing";

import { ContainerFactory } from "./ContainerFactory";
import { JobLogSourcePage } from './JobLogSourcePage';


/**
 * The site map node for this page.
 */
export const siteMapNode: ISiteMapNode = {
    caption: "Job Log Source",
    routeExpression: "/job_log_source/{id}",
    containerFactory: new ContainerFactory(),
    render: p => <JobLogSourcePage/>
};
