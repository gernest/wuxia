var _plan={
    title:"simple plan",
    template_engine:"go",
    dependency:[
        "markdown",
        "yaml_front_matter",
    ],
    strategies:[
        {
            title:"markdown",
            patterns:["*md"],
            exec:["markdown"],
        },
    ]
};

Wu.setCustomPlan(_plan);
