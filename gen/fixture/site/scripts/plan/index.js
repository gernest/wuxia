var _plan={
    title:"simple plan",
    template_engine:"go",
    dependencies:[
        "markdown",
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
