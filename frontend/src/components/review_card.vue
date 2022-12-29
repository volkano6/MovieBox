<template>
  <div class="post-list">
          
            <div class="row pt-3 rounded" style="background: #000;">
                <div class="col-sm-1">
                    <div class="picture">
                        <img alt="Opt wizard thumbnail" style="width: 70%; height: 70%; float: right;" src="https://bootdey.com/img/Content/user_2.jpg">
                    </div>
                </div>
                <div class="col-sm-8">
                    <h4>
                        <a :href="'/user/' + UserID" class="username">{{Username}}</a>
                        <a :href="'/movie/' + MovieID" class="filmname">
                          <label class="label label-info" style="color:cornflowerblue">&nbsp; #{{MovieTitle}} </label>
                        </a>
                        
                    </h4>
                    <h5> 
                    <i class="fa fa-calendar"></i>
                    {{CommentDate}}
                    </h5>
                    <p class="description" style="color: aliceblue;">{{Comment}}</p>    
                </div>
            </div>
        </div>
</template>

<script>
export default {
  props:["UserID", "Username","UserAvatar", "MovieTitle", "MovieID", "Comment", "CommentDate"] ,
}
</script>

<style scoped>
.post-list {
  position: relative;
  padding: 5px 0;
}

.post-list .picture {
  max-width: 400px;
  overflow: hidden;
  height: auto;
  border-radius: 6px;
}

.post-list .label{
    font-weight:normal;    
}

.post-list .picture {
  max-width: 210px;
}

.post-list .picture img {
  width: 100%;
}

.post-list h4 {
  font-size: 20px;
}

.post-list h5 {
  color: #888;
}

.post-list p {
  float: left;
}

.post-list:after {
  background: #EEEEEE;
  width: 83.3333%;
  bottom: 0;
  right: 0;
  content: "";
  display: block;
  position: absolute;
}

.post-list .btn-download {
  margin-top: 45px;
}

.btn-info {
  color: #1084FF;
  border-color: #269abc;
}

.btn-round {
  border-width: 1px;
  border-radius: 30px !important;
  opacity: 0.79;
  padding: 9px 18px;
}
.btn {
  border-width: 2px;
  background-color: rgba(0,0,0,0);
  font-weight: 400;
  opacity: 0.8;
  padding: 7px 16px;
}
</style>