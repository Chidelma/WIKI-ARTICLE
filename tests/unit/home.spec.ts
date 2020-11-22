import { mount } from "@vue/test-utils";
import Home from "@/views/Home.vue";

describe("Home", () => {

    const wrapper = mount(Home);

    test("article length before fetch", () => {

        const articles = wrapper.findAll('.article');

        expect(articles.length).toBe(0);
    });

    test("article length after fetch", () => {

        const articles = wrapper.findAll('.article');

        setTimeout(() => {
            expect(articles.length).toBeGreaterThan(0);
        }, 1000);
    });
});