def predict_proba(X: Iterable[str]):
        return np.array([predict_one_probas(tweet) for tweet in X])