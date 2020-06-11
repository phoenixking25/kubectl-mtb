<p>id: MTB-PL1-BC-HI-5
title: Block use of host IPC
benchmarkType: Behavioral Check
category: Host Isolation
description: Tenants should not be allowed to share the host&rsquo;s inter-process communication (IPC) namespace.
remediation: Define a <code>PodSecurityPolicy</code> with <code>hostIPC</code> set to <code>false</code> and map the policy to each tenant&rsquo;s namespace, or use a policy engine such as <a href="https://github.com/open-policy-agent/gatekeeper">OPA/Gatekeeper</a> or <a href="https://kyverno.io">Kyverno</a> to enforce that <code>hostPID</code> cannot be set to <code>true</code>.
profileLevel: 1</p>
