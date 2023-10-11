import {
	createBrowserRouter,
	createRoutesFromElements,
	Route,
} from "react-router-dom";

import App from "@/App";
import Home from "@/pages/Home";

const router = createBrowserRouter(
    createRoutesFromElements(
        <Route path="/" element={<App />}>
            <Route element={<Home />}>
                <Route index element={<div>Today</div>}/>
                <Route path="/incoming" element={<div>Incoming</div>}/>
            </Route>
        </Route>
    )
);

export default router;
