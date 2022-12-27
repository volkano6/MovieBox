<template>
  <div class="row mt-3 justify-content-center">
    <div class="col-9">
      <div v-for="value in comments" :key=value.id>
      <ReviewCard :MovieID=value.movie_id :Username=value.user_name :MovieTitle=value.movie_title :Comment=value.comment
        :CommentDate=value.comment_date />
    </div>
    </div>
  </div>
</template>

<script>
import ReviewCard from '../../review_card.vue';

export default {
name: "userComments",
props: ["comments"],
components: {
ReviewCard
}
}
</script>

<style>

</style>