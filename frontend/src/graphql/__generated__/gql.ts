/* eslint-disable */
import * as types from './graphql';
import { TypedDocumentNode as DocumentNode } from '@graphql-typed-document-node/core';

/**
 * Map of all GraphQL operations in the project.
 *
 * This map has several performance disadvantages:
 * 1. It is not tree-shakeable, so it will include all operations in the project.
 * 2. It is not minifiable, so the string of a GraphQL query will be multiple times inside the bundle.
 * 3. It does not support dead code elimination, so it will add unused operations.
 *
 * Therefore it is highly recommended to use the babel or swc plugin for production.
 */
const documents = {
    "\n  query LoadFullBazelInvocationDetails($invocationID: String!) {\n    bazelInvocation(invocationId: $invocationID) {\n      ...FullBazelInvocationDetails\n    }\n  }\n": types.LoadFullBazelInvocationDetailsDocument,
    "\n  fragment RunnerMetricsInfo on BazelInvocation {\n    id\n    metrics {\n      id\n      actionSummary {\n        id\n        runnerCount{\n            id\n            actionsExecuted\n            name\n            execKind\n        }\n      }\n    }\n  }\n": types.RunnerMetricsInfoFragmentDoc,
    "\nfragment BazelInvocationInfo on BazelInvocation {\n  metrics {\n    id\n    actionSummary {\n      id\n      actionsCreated\n      actionsExecuted\n      actionsCreatedNotIncludingAspects\n      remoteCacheHits\n      actionCacheStatistics {\n        id\n        loadTimeInMs\n        saveTimeInMs\n        hits\n        misses\n        sizeInBytes\n        missDetails {\n          id\n          count\n          reason\n        }\n      }\n      runnerCount {\n        id\n        actionsExecuted\n        name\n        execKind\n      }\n      actionData {\n        id\n        mnemonic\n        userTime\n        systemTime\n        lastEndedMs\n        actionsCreated\n        actionsExecuted\n        firstStartedMs\n      }\n    }\n    artifactMetrics {\n      id\n      sourceArtifactsRead {\n        id\n        sizeInBytes\n        count\n      }\n      outputArtifactsSeen {\n        id\n        sizeInBytes\n        count\n      }\n      outputArtifactsFromActionCache {\n        id\n        sizeInBytes\n        count\n      }\n      topLevelArtifacts {\n        id\n        sizeInBytes\n        count\n      }\n    }\n    cumulativeMetrics {\n      id\n      numBuilds\n      numAnalyses\n    }\n    dynamicExecutionMetrics {\n      id\n      raceStatistics {\n        id\n        localWins\n        mnemonic\n        renoteWins\n        localRunner\n        remoteRunner\n      }\n    }\n    buildGraphMetrics {\n      id\n      actionLookupValueCount\n      actionLookupValueCountNotIncludingAspects\n      actionCount\n      inputFileConfiguredTargetCount\n      outputFileConfiguredTargetCount\n      otherConfiguredTargetCount\n      outputArtifactCount\n      postInvocationSkyframeNodeCount\n    }\n    memoryMetrics {\n      id\n      usedHeapSizePostBuild\n      peakPostGcHeapSize\n      peakPostGcTenuredSpaceHeapSize\n      garbageMetrics {\n        id\n        garbageCollected\n        type\n      }\n    }\n    targetMetrics {\n      id\n      targetsLoaded\n      targetsConfigured\n      targetsConfiguredNotIncludingAspects\n    }\n    timingMetrics {\n      id\n      cpuTimeInMs\n      wallTimeInMs\n      analysisPhaseTimeInMs\n      executionPhaseTimeInMs\n      actionsExecutionStartInMs\n    }\n    networkMetrics {\n      id\n      systemNetworkStats {\n        id\n        bytesSent\n        bytesRecv\n        packetsSent\n        packetsRecv\n        peakBytesSentPerSec\n        peakBytesRecvPerSec\n        peakPacketsSentPerSec\n        peakPacketsRecvPerSec\n      }\n    }\n    packageMetrics {\n      id\n      packagesLoaded\n      packageLoadMetrics {\n        id\n        name\n        numTargets\n        loadDuration\n        packageOverhead\n        computationSteps\n        numTransitiveLoads\n      }\n    }\n  }\n  bazelCommand {\n    command\n    executable\n    id\n    buildOptions: options\n    residual\n  }\n  id\n  invocationID\n  build {\n    id\n    buildUUID\n  }\n\n  relatedFiles {\n    name\n    url\n  }\n  user {\n    Email\n    LDAP\n  }\n\n  startedAt\n  endedAt\n  state {\n    bepCompleted\n    buildEndTime\n    buildStartTime\n    exitCode {\n      code\n      id\n      name\n    }\n    id\n  }\n  stepLabel\n}\n": types.BazelInvocationInfoFragmentDoc,
    "\n fragment ProblemInfo on Problem {\n  id\n  label\n  __typename\n  ... on ActionProblem {\n    __typename\n    id\n    label\n    type\n    stdout {\n      ...BlobReferenceInfo\n    }\n    stderr {\n      ...BlobReferenceInfo\n    }\n  }\n  ... on TestProblem {\n    __typename\n    id\n    label\n    status\n    results {\n      __typename\n      id\n      run\n      shard\n      attempt\n      status\n      actionLogOutput {\n        ...BlobReferenceInfo\n      }\n      undeclaredTestOutputs {\n        ...BlobReferenceInfo\n      }\n    }\n  }\n  ... on TargetProblem {\n    __typename\n    id\n    label\n  }\n  ... on ProgressProblem {\n    __typename\n    id\n    output\n    label\n  }\n}\n": types.ProblemInfoFragmentDoc,
    "\nfragment BlobReferenceInfo on BlobReference {\n  availabilityStatus\n  name\n  sizeInBytes\n  downloadURL\n}\n": types.BlobReferenceInfoFragmentDoc,
    "\n    fragment FullBazelInvocationDetails on BazelInvocation {\n      problems {\n        ...ProblemInfo\n      }\n      ...BazelInvocationInfo\n    }\n": types.FullBazelInvocationDetailsFragmentDoc,
    "\n  query GetActionProblem($id: ID!) {\n    node(id: $id) {\n      id\n      ... on ActionProblem {\n        label\n        stdout {\n          ...BlobReferenceInfo\n        }\n        stderr {\n          ...BlobReferenceInfo\n        }\n      }\n    }\n  }\n": types.GetActionProblemDocument,
    "\nfragment TestResultInfo on TestResult {\n      actionLogOutput {\n  ...BlobReferenceInfo\n  }\n  attempt\n  run\n  shard\n  status\n  undeclaredTestOutputs {\n    ...BlobReferenceInfo\n  }\n}": types.TestResultInfoFragmentDoc,
    "\n  query FindBuildByUUID($url: String, $uuid: UUID) {\n    getBuild(buildURL: $url, buildUUID: $uuid) {\n      id\n      buildURL\n      buildUUID\n      invocations {\n        ...FullBazelInvocationDetails\n      }\n      env {\n        key\n        value\n      }\n    }\n  }\n": types.FindBuildByUuidDocument,
    "\n  query FindBazelInvocations(\n    $first: Int!\n    $where: BazelInvocationWhereInput\n  ) {\n    findBazelInvocations(first: $first, where: $where) {\n      edges {\n        node {\n          ...BazelInvocationNode\n        }\n      }\n    }\n  }\n": types.FindBazelInvocationsDocument,
    "\n  fragment BazelInvocationNode on BazelInvocation {\n    id\n    invocationID\n    startedAt\n    user {\n      Email\n      LDAP\n    }\n    endedAt\n    state {\n      bepCompleted\n      exitCode {\n        code\n        name\n      }\n    }\n    build {\n      buildUUID\n    }\n  }\n": types.BazelInvocationNodeFragmentDoc,
    "\n  query FindBuilds(\n    $first: Int!\n    $where: BuildWhereInput\n  ) {\n    findBuilds(first: $first, where: $where) {\n      edges {\n        node {\n          ...BuildNode\n        }\n      }\n    }\n  }\n": types.FindBuildsDocument,
    "\n  fragment BuildNode on Build {\n    id\n    buildUUID\n    buildURL\n  }\n": types.BuildNodeFragmentDoc,
};

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 *
 *
 * @example
 * ```ts
 * const query = gql(`query GetUser($id: ID!) { user(id: $id) { name } }`);
 * ```
 *
 * The query argument is unknown!
 * Please regenerate the types.
 */
export function gql(source: string): unknown;

/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query LoadFullBazelInvocationDetails($invocationID: String!) {\n    bazelInvocation(invocationId: $invocationID) {\n      ...FullBazelInvocationDetails\n    }\n  }\n"): (typeof documents)["\n  query LoadFullBazelInvocationDetails($invocationID: String!) {\n    bazelInvocation(invocationId: $invocationID) {\n      ...FullBazelInvocationDetails\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment RunnerMetricsInfo on BazelInvocation {\n    id\n    metrics {\n      id\n      actionSummary {\n        id\n        runnerCount{\n            id\n            actionsExecuted\n            name\n            execKind\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  fragment RunnerMetricsInfo on BazelInvocation {\n    id\n    metrics {\n      id\n      actionSummary {\n        id\n        runnerCount{\n            id\n            actionsExecuted\n            name\n            execKind\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nfragment BazelInvocationInfo on BazelInvocation {\n  metrics {\n    id\n    actionSummary {\n      id\n      actionsCreated\n      actionsExecuted\n      actionsCreatedNotIncludingAspects\n      remoteCacheHits\n      actionCacheStatistics {\n        id\n        loadTimeInMs\n        saveTimeInMs\n        hits\n        misses\n        sizeInBytes\n        missDetails {\n          id\n          count\n          reason\n        }\n      }\n      runnerCount {\n        id\n        actionsExecuted\n        name\n        execKind\n      }\n      actionData {\n        id\n        mnemonic\n        userTime\n        systemTime\n        lastEndedMs\n        actionsCreated\n        actionsExecuted\n        firstStartedMs\n      }\n    }\n    artifactMetrics {\n      id\n      sourceArtifactsRead {\n        id\n        sizeInBytes\n        count\n      }\n      outputArtifactsSeen {\n        id\n        sizeInBytes\n        count\n      }\n      outputArtifactsFromActionCache {\n        id\n        sizeInBytes\n        count\n      }\n      topLevelArtifacts {\n        id\n        sizeInBytes\n        count\n      }\n    }\n    cumulativeMetrics {\n      id\n      numBuilds\n      numAnalyses\n    }\n    dynamicExecutionMetrics {\n      id\n      raceStatistics {\n        id\n        localWins\n        mnemonic\n        renoteWins\n        localRunner\n        remoteRunner\n      }\n    }\n    buildGraphMetrics {\n      id\n      actionLookupValueCount\n      actionLookupValueCountNotIncludingAspects\n      actionCount\n      inputFileConfiguredTargetCount\n      outputFileConfiguredTargetCount\n      otherConfiguredTargetCount\n      outputArtifactCount\n      postInvocationSkyframeNodeCount\n    }\n    memoryMetrics {\n      id\n      usedHeapSizePostBuild\n      peakPostGcHeapSize\n      peakPostGcTenuredSpaceHeapSize\n      garbageMetrics {\n        id\n        garbageCollected\n        type\n      }\n    }\n    targetMetrics {\n      id\n      targetsLoaded\n      targetsConfigured\n      targetsConfiguredNotIncludingAspects\n    }\n    timingMetrics {\n      id\n      cpuTimeInMs\n      wallTimeInMs\n      analysisPhaseTimeInMs\n      executionPhaseTimeInMs\n      actionsExecutionStartInMs\n    }\n    networkMetrics {\n      id\n      systemNetworkStats {\n        id\n        bytesSent\n        bytesRecv\n        packetsSent\n        packetsRecv\n        peakBytesSentPerSec\n        peakBytesRecvPerSec\n        peakPacketsSentPerSec\n        peakPacketsRecvPerSec\n      }\n    }\n    packageMetrics {\n      id\n      packagesLoaded\n      packageLoadMetrics {\n        id\n        name\n        numTargets\n        loadDuration\n        packageOverhead\n        computationSteps\n        numTransitiveLoads\n      }\n    }\n  }\n  bazelCommand {\n    command\n    executable\n    id\n    buildOptions: options\n    residual\n  }\n  id\n  invocationID\n  build {\n    id\n    buildUUID\n  }\n\n  relatedFiles {\n    name\n    url\n  }\n  user {\n    Email\n    LDAP\n  }\n\n  startedAt\n  endedAt\n  state {\n    bepCompleted\n    buildEndTime\n    buildStartTime\n    exitCode {\n      code\n      id\n      name\n    }\n    id\n  }\n  stepLabel\n}\n"): (typeof documents)["\nfragment BazelInvocationInfo on BazelInvocation {\n  metrics {\n    id\n    actionSummary {\n      id\n      actionsCreated\n      actionsExecuted\n      actionsCreatedNotIncludingAspects\n      remoteCacheHits\n      actionCacheStatistics {\n        id\n        loadTimeInMs\n        saveTimeInMs\n        hits\n        misses\n        sizeInBytes\n        missDetails {\n          id\n          count\n          reason\n        }\n      }\n      runnerCount {\n        id\n        actionsExecuted\n        name\n        execKind\n      }\n      actionData {\n        id\n        mnemonic\n        userTime\n        systemTime\n        lastEndedMs\n        actionsCreated\n        actionsExecuted\n        firstStartedMs\n      }\n    }\n    artifactMetrics {\n      id\n      sourceArtifactsRead {\n        id\n        sizeInBytes\n        count\n      }\n      outputArtifactsSeen {\n        id\n        sizeInBytes\n        count\n      }\n      outputArtifactsFromActionCache {\n        id\n        sizeInBytes\n        count\n      }\n      topLevelArtifacts {\n        id\n        sizeInBytes\n        count\n      }\n    }\n    cumulativeMetrics {\n      id\n      numBuilds\n      numAnalyses\n    }\n    dynamicExecutionMetrics {\n      id\n      raceStatistics {\n        id\n        localWins\n        mnemonic\n        renoteWins\n        localRunner\n        remoteRunner\n      }\n    }\n    buildGraphMetrics {\n      id\n      actionLookupValueCount\n      actionLookupValueCountNotIncludingAspects\n      actionCount\n      inputFileConfiguredTargetCount\n      outputFileConfiguredTargetCount\n      otherConfiguredTargetCount\n      outputArtifactCount\n      postInvocationSkyframeNodeCount\n    }\n    memoryMetrics {\n      id\n      usedHeapSizePostBuild\n      peakPostGcHeapSize\n      peakPostGcTenuredSpaceHeapSize\n      garbageMetrics {\n        id\n        garbageCollected\n        type\n      }\n    }\n    targetMetrics {\n      id\n      targetsLoaded\n      targetsConfigured\n      targetsConfiguredNotIncludingAspects\n    }\n    timingMetrics {\n      id\n      cpuTimeInMs\n      wallTimeInMs\n      analysisPhaseTimeInMs\n      executionPhaseTimeInMs\n      actionsExecutionStartInMs\n    }\n    networkMetrics {\n      id\n      systemNetworkStats {\n        id\n        bytesSent\n        bytesRecv\n        packetsSent\n        packetsRecv\n        peakBytesSentPerSec\n        peakBytesRecvPerSec\n        peakPacketsSentPerSec\n        peakPacketsRecvPerSec\n      }\n    }\n    packageMetrics {\n      id\n      packagesLoaded\n      packageLoadMetrics {\n        id\n        name\n        numTargets\n        loadDuration\n        packageOverhead\n        computationSteps\n        numTransitiveLoads\n      }\n    }\n  }\n  bazelCommand {\n    command\n    executable\n    id\n    buildOptions: options\n    residual\n  }\n  id\n  invocationID\n  build {\n    id\n    buildUUID\n  }\n\n  relatedFiles {\n    name\n    url\n  }\n  user {\n    Email\n    LDAP\n  }\n\n  startedAt\n  endedAt\n  state {\n    bepCompleted\n    buildEndTime\n    buildStartTime\n    exitCode {\n      code\n      id\n      name\n    }\n    id\n  }\n  stepLabel\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n fragment ProblemInfo on Problem {\n  id\n  label\n  __typename\n  ... on ActionProblem {\n    __typename\n    id\n    label\n    type\n    stdout {\n      ...BlobReferenceInfo\n    }\n    stderr {\n      ...BlobReferenceInfo\n    }\n  }\n  ... on TestProblem {\n    __typename\n    id\n    label\n    status\n    results {\n      __typename\n      id\n      run\n      shard\n      attempt\n      status\n      actionLogOutput {\n        ...BlobReferenceInfo\n      }\n      undeclaredTestOutputs {\n        ...BlobReferenceInfo\n      }\n    }\n  }\n  ... on TargetProblem {\n    __typename\n    id\n    label\n  }\n  ... on ProgressProblem {\n    __typename\n    id\n    output\n    label\n  }\n}\n"): (typeof documents)["\n fragment ProblemInfo on Problem {\n  id\n  label\n  __typename\n  ... on ActionProblem {\n    __typename\n    id\n    label\n    type\n    stdout {\n      ...BlobReferenceInfo\n    }\n    stderr {\n      ...BlobReferenceInfo\n    }\n  }\n  ... on TestProblem {\n    __typename\n    id\n    label\n    status\n    results {\n      __typename\n      id\n      run\n      shard\n      attempt\n      status\n      actionLogOutput {\n        ...BlobReferenceInfo\n      }\n      undeclaredTestOutputs {\n        ...BlobReferenceInfo\n      }\n    }\n  }\n  ... on TargetProblem {\n    __typename\n    id\n    label\n  }\n  ... on ProgressProblem {\n    __typename\n    id\n    output\n    label\n  }\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nfragment BlobReferenceInfo on BlobReference {\n  availabilityStatus\n  name\n  sizeInBytes\n  downloadURL\n}\n"): (typeof documents)["\nfragment BlobReferenceInfo on BlobReference {\n  availabilityStatus\n  name\n  sizeInBytes\n  downloadURL\n}\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n    fragment FullBazelInvocationDetails on BazelInvocation {\n      problems {\n        ...ProblemInfo\n      }\n      ...BazelInvocationInfo\n    }\n"): (typeof documents)["\n    fragment FullBazelInvocationDetails on BazelInvocation {\n      problems {\n        ...ProblemInfo\n      }\n      ...BazelInvocationInfo\n    }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query GetActionProblem($id: ID!) {\n    node(id: $id) {\n      id\n      ... on ActionProblem {\n        label\n        stdout {\n          ...BlobReferenceInfo\n        }\n        stderr {\n          ...BlobReferenceInfo\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query GetActionProblem($id: ID!) {\n    node(id: $id) {\n      id\n      ... on ActionProblem {\n        label\n        stdout {\n          ...BlobReferenceInfo\n        }\n        stderr {\n          ...BlobReferenceInfo\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\nfragment TestResultInfo on TestResult {\n      actionLogOutput {\n  ...BlobReferenceInfo\n  }\n  attempt\n  run\n  shard\n  status\n  undeclaredTestOutputs {\n    ...BlobReferenceInfo\n  }\n}"): (typeof documents)["\nfragment TestResultInfo on TestResult {\n      actionLogOutput {\n  ...BlobReferenceInfo\n  }\n  attempt\n  run\n  shard\n  status\n  undeclaredTestOutputs {\n    ...BlobReferenceInfo\n  }\n}"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query FindBuildByUUID($url: String, $uuid: UUID) {\n    getBuild(buildURL: $url, buildUUID: $uuid) {\n      id\n      buildURL\n      buildUUID\n      invocations {\n        ...FullBazelInvocationDetails\n      }\n      env {\n        key\n        value\n      }\n    }\n  }\n"): (typeof documents)["\n  query FindBuildByUUID($url: String, $uuid: UUID) {\n    getBuild(buildURL: $url, buildUUID: $uuid) {\n      id\n      buildURL\n      buildUUID\n      invocations {\n        ...FullBazelInvocationDetails\n      }\n      env {\n        key\n        value\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query FindBazelInvocations(\n    $first: Int!\n    $where: BazelInvocationWhereInput\n  ) {\n    findBazelInvocations(first: $first, where: $where) {\n      edges {\n        node {\n          ...BazelInvocationNode\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query FindBazelInvocations(\n    $first: Int!\n    $where: BazelInvocationWhereInput\n  ) {\n    findBazelInvocations(first: $first, where: $where) {\n      edges {\n        node {\n          ...BazelInvocationNode\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment BazelInvocationNode on BazelInvocation {\n    id\n    invocationID\n    startedAt\n    user {\n      Email\n      LDAP\n    }\n    endedAt\n    state {\n      bepCompleted\n      exitCode {\n        code\n        name\n      }\n    }\n    build {\n      buildUUID\n    }\n  }\n"): (typeof documents)["\n  fragment BazelInvocationNode on BazelInvocation {\n    id\n    invocationID\n    startedAt\n    user {\n      Email\n      LDAP\n    }\n    endedAt\n    state {\n      bepCompleted\n      exitCode {\n        code\n        name\n      }\n    }\n    build {\n      buildUUID\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  query FindBuilds(\n    $first: Int!\n    $where: BuildWhereInput\n  ) {\n    findBuilds(first: $first, where: $where) {\n      edges {\n        node {\n          ...BuildNode\n        }\n      }\n    }\n  }\n"): (typeof documents)["\n  query FindBuilds(\n    $first: Int!\n    $where: BuildWhereInput\n  ) {\n    findBuilds(first: $first, where: $where) {\n      edges {\n        node {\n          ...BuildNode\n        }\n      }\n    }\n  }\n"];
/**
 * The gql function is used to parse GraphQL queries into a document that can be used by GraphQL clients.
 */
export function gql(source: "\n  fragment BuildNode on Build {\n    id\n    buildUUID\n    buildURL\n  }\n"): (typeof documents)["\n  fragment BuildNode on Build {\n    id\n    buildUUID\n    buildURL\n  }\n"];

export function gql(source: string) {
  return (documents as any)[source] ?? {};
}

export type DocumentType<TDocumentNode extends DocumentNode<any, any>> = TDocumentNode extends DocumentNode<  infer TType,  any>  ? TType  : never;