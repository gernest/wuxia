var _plan={};
_plan.title="simple plan";
_plan.dependency=[
  "markdown",
  "yaml_front_matter",
];
_plan.template_engine="go";

_post_strategy={};
_post_strategy.pattern='*.md';
_post_strategy.exec=["markdown"];
_plan.strategies=[_post_strategy];

// Register the plan
System.plan=_plan; 

