package internal

import (
	"time"
)

const refType = "github"

type pageInfo struct {
	HasNextPage bool   `json:"hasNextPage"`
	EndCursor   string `json:"endCursor"`
}

type rateLimit struct {
	Limit     int       `json:"limit"`
	Cost      int       `json:"cost"`
	Remaining int       `json:"remaining"`
	ResetAt   time.Time `json:"resetAt"`
}

func (l rateLimit) ShouldPause() bool {
	// stop at 80%
	return float32(l.Remaining)*.8 >= float32(l.Limit)
}

type nameProp struct {
	Name string `json:"name"`
}

type oidProp struct {
	Oid string `json:"oid"`
}

type organization struct {
	Repositories repositories `json:"repositories"`
}

type allQueryResult struct {
	Organization organization `json:"organization"`
	RateLimit    rateLimit    `json:"rateLimit"`
}

var pullrequestPageQuery = `
query GetPullRequests($name: String!, $owner: String!, $first: Int!, $after: String) {
	repository(name: $name, owner: $owner) {
		pullRequests(first: $first, after: $after, orderBy: {field: UPDATED_AT, direction: DESC}) {
			totalCount
			pageInfo {
				hasNextPage
				endCursor
			}
			nodes {
				id
				bodyHTML
				url
				closed
				draft: isDraft
				locked
				merged
				number
				state
				title
				createdAt
				updatedAt
				mergedAt
				branch: headRefName
				mergeCommit {
					oid
				}
				mergedBy {
					avatarUrl
					...on User {
						id
						login
						email
						name
					}
					...on Bot {
						boturl: url
					}
				}
				author {
					avatarUrl
					...on User {
						id
						login
						email
						name
					}
					...on Bot {
						boturl: url
					}
				}
				commits(first: 10) {
					totalCount
					pageInfo {
						hasNextPage
						endCursor
					}
					nodes {
						commit {
							sha: oid
							message
							authoredDate
							additions
							deletions
							url
							author {
								avatarUrl
								email
								name
								user {
									id
									login
								}
							}
							committer {
								avatarUrl
								email
								name
								user {
									id
									login
								}
							}
						}
					}
				}
				reviews(first: 10) {
					nodes {
						id
						state
						createdAt
						url
						author {
							avatarUrl
							...on User {
								id
								login
								email
								name
							}
							...on Bot {
								boturl: url
							}
						}
					}
				}
			}
		}
	}
	rateLimit {
		limit
		cost
		remaining
		resetAt
	}
}
`

type allOrgViewOrg struct {
	Organizations organizations `json:"organizations"`
}

type allOrgsResult struct {
	Viewer    allOrgViewOrg `json:"viewer"`
	RateLimit rateLimit     `json:"rateLimit"`
}

type org struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	IsMember bool   `json:"viewerIsAMember"`
	IsAdmin  bool   `json:"viewerCanAdminister"`
}

type organizations struct {
	Nodes []org `json:"nodes"`
}

var allOrgsQuery = `
query GetAllOrgs($first: Int!, $after: String) {
	viewer {
		organizations(first: $first after: $after) {
			nodes {
				name
				login
				viewerIsAMember
				viewerCanAdminister
			}
		}
	}
}
`

var allDataQuery = `
query GetAllData($login: String!, $first: Int!, $after: String) {
	organization(login: $login) {
		repositories(first: $first, after: $after, isFork: false, orderBy: {field: UPDATED_AT, direction: DESC}) {
			totalCount
			pageInfo {
				hasNextPage
				endCursor
			}
			nodes {
				id
				nameWithOwner
				url
				updatedAt
				description
				defaultBranchRef {
					name
				}
				primaryLanguage {
					name
				}
				isArchived

				pullRequests(first: 10, orderBy: {field: UPDATED_AT, direction: DESC}) {
					totalCount
					pageInfo {
						hasNextPage
						endCursor
					}
					nodes {
						id
						bodyHTML
						url
						closed
						draft: isDraft
						locked
						merged
						number
						state
						title
						createdAt
						updatedAt
						mergedAt
						branch: headRefName
						mergeCommit {
							oid
						}
						mergedBy {
							avatarUrl
							...on User {
								id
								login
								email
								name
							}
							...on Bot {
								boturl: url
							}
						}
						author {
							avatarUrl
							...on User {
								id
								login
								email
								name
							}
							...on Bot {
								boturl: url
							}
						}
						commits(first: 1) {
							totalCount
							pageInfo {
								hasNextPage
								endCursor
							}
							nodes {
								commit {
									sha: oid
									message
									authoredDate
									additions
									deletions
									url
									author {
										avatarUrl
										email
										name
										user {
											id
											login
										}
									}
									committer {
										avatarUrl
										email
										name
										user {
											id
											login
										}
									}
								}
							}
						}
						reviews(first: 10) {
							nodes {
								id
								state
								createdAt
								url
								author {
									avatarUrl
									...on User {
										id
										login
										email
										name
									}
									...on Bot {
										boturl: url
									}
								}
							}
						}
					}
				}
			}
		}
	}
	rateLimit {
		limit
		cost
		remaining
		resetAt
	}
}
`
